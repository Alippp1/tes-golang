package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"github.com/Alippp1/tes-golang/internal/dto"
	"github.com/Alippp1/tes-golang/internal/models"
	"github.com/Alippp1/tes-golang/internal/repository"
)

type PurchasingService interface {
	Create(userID uint, req dto.CreatePurchasingRequest) (*dto.PurchasingResponse, error)
	FindAll() ([]dto.PurchasingResponse, error)
	FindByID(id uint) (*dto.PurchasingResponse, error)
}

type purchasingService struct {
	db           *gorm.DB
	purchaseRepo repository.PurchasingRepository
	detailRepo   repository.PurchasingDetailRepository
	itemRepo     repository.ItemRepository
}

func NewPurchasingService(
	db *gorm.DB,
	purchaseRepo repository.PurchasingRepository,
	detailRepo repository.PurchasingDetailRepository,
	itemRepo repository.ItemRepository,
) PurchasingService {
	return &purchasingService{
		db, purchaseRepo, detailRepo, itemRepo,
	}
}

func (s *purchasingService) Create(userID uint, req dto.CreatePurchasingRequest) (*dto.PurchasingResponse, error) {
	var (
		grandTotal float64
		details    []models.PurchasingDetail
		respDetail []dto.PurchasingDetailResponse
	)

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	purchasing := models.Purchasing{
		Date:       req.Date,
		SupplierID: req.SupplierID,
		UserID:     userID,
	}

	for _, itemReq := range req.Items {
		item, err := s.itemRepo.FindByID(itemReq.ItemID)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("item not found")
		}

		subTotal := float64(itemReq.Qty) * item.Price
		grandTotal += subTotal

		// Update stock
		item.Stock += itemReq.Qty
		if err := s.itemRepo.Updatetx(tx, item); err != nil {
			tx.Rollback()
			return nil, err
		}

		details = append(details, models.PurchasingDetail{
			ItemID:   item.ID,
			Qty:      itemReq.Qty,
			SubTotal: subTotal,
		})

		respDetail = append(respDetail, dto.PurchasingDetailResponse{
			ItemID:   item.ID,
			ItemName: item.Name,
			Qty:      itemReq.Qty,
			Price:    item.Price,
			SubTotal: subTotal,
		})
	}

	purchasing.GrandTotal = grandTotal

	if err := s.purchaseRepo.Create(tx, &purchasing); err != nil {
		tx.Rollback()
		return nil, err
	}

	for i := range details {
		details[i].PurchasingID = purchasing.ID
	}

	if err := s.detailRepo.BulkCreate(tx, details); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// 🚀 Webhook (bonus)
	go sendWebhook(purchasing, respDetail)

	return &dto.PurchasingResponse{
		ID:         purchasing.ID,
		Date:       purchasing.Date,
		SupplierID: purchasing.SupplierID,
		GrandTotal: purchasing.GrandTotal,
		Details:    respDetail,
	}, nil
}

func sendWebhook(p models.Purchasing, details []dto.PurchasingDetailResponse) {
	payload := map[string]interface{}{
		"purchasing_id": p.ID,
		"date":          p.Date,
		"supplier_id":   p.SupplierID,
		"grand_total":   p.GrandTotal,
		"details":       details,
	}

	body, _ := json.Marshal(payload)

	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Println("WEBHOOK_URL not set")
		return
	}

	_, err := http.Post(
		webhookURL,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		log.Println("webhook error:", err)
	}
}

func (s *purchasingService) FindAll() ([]dto.PurchasingResponse, error) {
	purchases, err := s.purchaseRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var resp []dto.PurchasingResponse

	for _, p := range purchases {
		var details []dto.PurchasingDetailResponse

		for _, d := range p.Details {
			details = append(details, dto.PurchasingDetailResponse{
				ItemID:   d.ItemID,
				ItemName: d.Item.Name,
				Qty:      d.Qty,
				Price:    d.Item.Price,
				SubTotal: d.SubTotal,
			})
		}

		resp = append(resp, dto.PurchasingResponse{
			ID:         p.ID,
			Date:       p.Date,
			SupplierID: p.SupplierID,
			Supplier:   p.Supplier.Name,
			UserID:     p.UserID,
			User:       p.User.Username,
			GrandTotal: p.GrandTotal,
			Details:    details,
		})
	}

	return resp, nil
}

func (s *purchasingService) FindByID(id uint) (*dto.PurchasingResponse, error) {
	p, err := s.purchaseRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("purchasing not found")
		}
		return nil, err
	}

	var details []dto.PurchasingDetailResponse
	for _, d := range p.Details {
		details = append(details, dto.PurchasingDetailResponse{
			ItemID:   d.ItemID,
			ItemName: d.Item.Name,
			Qty:      d.Qty,
			Price:    d.Item.Price,
			SubTotal: d.SubTotal,
		})
	}

	return &dto.PurchasingResponse{
		ID:         p.ID,
		Date:       p.Date,
		SupplierID: p.SupplierID,
		Supplier:   p.Supplier.Name,
		UserID:     p.UserID,
		User:       p.User.Username,
		GrandTotal: p.GrandTotal,
		Details:    details,
	}, nil
}
