package service

import (
	"errors"
	"log"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/model"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/repository"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
)

type EventService struct {
    Repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
    return &EventService{Repo: repo}
}

func parseDate(dateStr string) (*time.Time, error) {
    if dateStr == "" {
        return nil, nil
    }
    t, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return nil, err
    }
    return &t, nil
}

func extractStringSlice(value interface{}) []string {
    var result []string
    
    switch v := value.(type) {
    case []string:
        return v
    case []interface{}:
        for _, item := range v {
            if str, ok := item.(string); ok && str != "" {
                result = append(result, str)
            }
        }
    case string:
        if v != "" {
            result = append(result, v)
        }
    case []byte:
        if len(v) > 0 {
            result = append(result, string(v))
        }
    }
    
    return result
}

func (s *EventService) Create(req map[string]interface{}, file multipart.File, header *multipart.FileHeader) (*model.Event, error) {
    log.Printf("Creating new event with request keys: %v", req)
    
    title := getString(req, "title")
    if title == "" {
        return nil, errors.New("title is required")
    }
    
    category := getString(req, "category")
    if category == "" {
        return nil, errors.New("category is required")
    }
    
    id := getString(req, "id")
    if id == "" {
        id = uuid.New().String()
        log.Printf("Generated new event ID: %s", id)
    }
    
    dateStr := getString(req, "date")
    date, err := parseDate(dateStr)
    if err != nil {
        log.Printf("Warning: date parse error: %v", err)
    }
    
    posterPath := ""
    if file != nil && header != nil {
        posterPath, err = util.SaveMultipartFile("storage/posters", file, header)
        if err != nil {
            log.Printf("Poster save error: %v", err)
            return nil, err
        }
        log.Printf("Saved poster to: %s", posterPath)
    }
    
    rules := extractStringSlice(req["rules"])
    log.Printf("Extracted rules: %v", rules)
    
    benefits := extractStringSlice(req["benefits"])
    log.Printf("Extracted benefits: %v", benefits)
    
    quota := 0
    switch q := req["quota"].(type) {
    case float64:
        quota = int(q)
    case int:
        quota = q
    case string:
        if q != "" {
            if qi, err := strconv.Atoi(q); err == nil {
                quota = qi
            }
        }
    case []byte:
        if len(q) > 0 {
            if qi, err := strconv.Atoi(string(q)); err == nil {
                quota = qi
            }
        }
    }
    log.Printf("Parsed quota: %d", quota)

    // Parse event_type and price
    eventType := getString(req, "eventType")
    if eventType == "" {
        eventType = "free"
    }
    log.Printf("Parsed event_type: %s", eventType)
    
    price := 0
    if eventType == "paid" {
        switch p := req["price"].(type) {
        case float64:
            price = int(p)
        case int:
            price = p
        case string:
            if p != "" {
                if pi, err := strconv.Atoi(p); err == nil {
                    price = pi
                }
            }
        case []byte:
            if len(p) > 0 {
                if pi, err := strconv.Atoi(string(p)); err == nil {
                    price = pi
                }
            }
        }
    }
    log.Printf("Parsed price: %d", price)

    available := true
    if av, ok := req["available"].(bool); ok {
        available = av
    } else if avs, ok := req["available"].(string); ok {
        available = avs == "1" || avs == "true" || avs == "on"
    }
    
    ev := &model.Event{
        ID:            id,
        Category:      category,
        CategoryLabel: getString(req, "categoryLabel"),
        Title:         title,
        Description:   getString(req, "description"),
        Rules:         rules,
        Benefits:      benefits,
        Date:          date,
        Time:          getString(req, "time"),
        Speaker1:      getString(req, "speaker1"),
        Speaker2:      getString(req, "speaker2"),
        Speaker3:      getString(req, "speaker3"),
        Moderator:     getString(req, "moderator"),
        Location:      getString(req, "location"),
        Quota:         quota,
        Registered:    0,
        Poster:        posterPath,
        EventType:     eventType,
        Price:         price,
        Available:     available,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }
    
    if ev.Date != nil {
        if isDatePassed(ev.Date) {
            ev.Available = false
            log.Printf("Event date already passed; setting Available=false")
        } else {
            ev.Available = true
            log.Printf("Event date is upcoming; setting Available=true")
        }
    }
    log.Printf("Creating event: ID=%s, Title=%s, Category=%s, EventType=%s, Price=%d", ev.ID, ev.Title, ev.Category, ev.EventType, ev.Price)
    
    if err := s.Repo.Create(ev); err != nil {
        log.Printf("Repository create error: %v", err)
        return nil, err
    }
    
    log.Printf("Event created successfully: %s", ev.ID)
    return ev, nil
}

func (s *EventService) Update(id string, req map[string]interface{}, file multipart.File, header *multipart.FileHeader) (*model.Event, error) {
    log.Printf("=== Starting update for event ID: %s ===", id)
    log.Printf("Request data: %+v", req)
    
    ev, err := s.Repo.GetByID(id)
    if err != nil {
        log.Printf("Event not found: %v", err)
        return nil, errors.New("event not found")
    }
    
    log.Printf("Found existing event: %s - %s", ev.ID, ev.Title)
    
    if v, ok := req["title"].(string); ok && v != "" {
        ev.Title = v
        log.Printf("✓ Updated title: %s", v)
    }
    
    if v, ok := req["description"].(string); ok {
        ev.Description = v
        log.Printf("✓ Updated description")
    }
    
    if v, ok := req["category"].(string); ok && v != "" {
        ev.Category = v
        log.Printf("✓ Updated category: %s", v)
    }
    
    if v, ok := req["categoryLabel"].(string); ok && v != "" {
        ev.CategoryLabel = v
        log.Printf("✓ Updated categoryLabel: %s", v)
    }
    
    if dateStr, ok := req["date"].(string); ok {
        if d, err := parseDate(dateStr); err == nil {
            ev.Date = d
            log.Printf("✓ Updated date: %s", dateStr)
        } else {
            log.Printf("⚠ Date parse error: %v", err)
        }
    }
    
    if v, ok := req["time"].(string); ok {
        ev.Time = v
        log.Printf("✓ Updated time: %s", v)
    }
    
    if v, ok := req["speaker1"].(string); ok {
        ev.Speaker1 = v
        log.Printf("✓ Updated speaker1: %s", v)
    }
    
    if v, ok := req["speaker2"].(string); ok {
        ev.Speaker2 = v
        log.Printf("✓ Updated speaker2: %s", v)
    }
    
    if v, ok := req["speaker3"].(string); ok {
        ev.Speaker3 = v
        log.Printf("✓ Updated speaker3: %s", v)
    }
    
    if v, ok := req["moderator"].(string); ok {
        ev.Moderator = v
        log.Printf("✓ Updated moderator: %s", v)
    }
    
    if v, ok := req["location"].(string); ok {
        ev.Location = v
        log.Printf("✓ Updated location: %s", v)
    }
    
    if _, hasQuota := req["quota"]; hasQuota {
        switch q := req["quota"].(type) {
        case float64:
            ev.Quota = int(q)
            log.Printf("✓ Updated quota (float): %d", ev.Quota)
        case int:
            ev.Quota = q
            log.Printf("✓ Updated quota (int): %d", ev.Quota)
        case string:
            if q != "" {
                if qi, err := strconv.Atoi(q); err == nil {
                    ev.Quota = qi
                    log.Printf("✓ Updated quota (string): %d", ev.Quota)
                }
            }
        case []byte:
            if len(q) > 0 {
                if qi, err := strconv.Atoi(string(q)); err == nil {
                    ev.Quota = qi
                    log.Printf("✓ Updated quota (bytes): %d", ev.Quota)
                }
            }
        }
    }
    
    // Update event_type
    if v, ok := req["eventType"].(string); ok && v != "" {
        ev.EventType = v
        log.Printf("✓ Updated eventType: %s", v)
    }
    
    // Update price (only if event_type is paid)
    if ev.EventType == "paid" {
        if _, hasPrice := req["price"]; hasPrice {
            switch p := req["price"].(type) {
            case float64:
                ev.Price = int(p)
                log.Printf("✓ Updated price (float): %d", ev.Price)
            case int:
                ev.Price = p
                log.Printf("✓ Updated price (int): %d", ev.Price)
            case string:
                if p != "" {
                    if pi, err := strconv.Atoi(p); err == nil {
                        ev.Price = pi
                        log.Printf("✓ Updated price (string): %d", ev.Price)
                    }
                }
            case []byte:
                if len(p) > 0 {
                    if pi, err := strconv.Atoi(string(p)); err == nil {
                        ev.Price = pi
                        log.Printf("✓ Updated price (bytes): %d", ev.Price)
                    }
                }
            }
        }
    } else {
        ev.Price = 0
        log.Printf("✓ Event is free, setting price to 0")
    }
    
    if _, hasAvail := req["available"]; hasAvail {
        if av, ok := req["available"].(bool); ok {
            ev.Available = av
            log.Printf("✓ Updated available: %v", av)
        } else if avs, ok := req["available"].(string); ok {
            newAvail := avs == "1" || avs == "true" || avs == "on"
            ev.Available = newAvail
            log.Printf("✓ Updated available from string: %v", newAvail)
        }
    }
    
    if _, hasRules := req["rules"]; hasRules {
        rules := extractStringSlice(req["rules"])
        ev.Rules = rules
        log.Printf("✓ Updated rules: %v", rules)
    } else {
        log.Printf("⚠ Rules not provided, keeping existing: %v", ev.Rules)
    }
    
    if _, hasBenefits := req["benefits"]; hasBenefits {
        benefits := extractStringSlice(req["benefits"])
        ev.Benefits = benefits
        log.Printf("✓ Updated benefits: %v", benefits)
    } else {
        log.Printf("⚠ Benefits not provided, keeping existing: %v", ev.Benefits)
    }
    
    if file != nil && header != nil {
        posterPath, err := util.SaveMultipartFile("storage/posters", file, header)
        if err != nil {
            log.Printf("Poster save error: %v", err)
            return nil, err
        }
        ev.Poster = posterPath
        log.Printf("✓ Updated poster: %s", posterPath)
    }
    
    ev.UpdatedAt = time.Now()

    if ev.Date != nil {
        if isDatePassed(ev.Date) {
            ev.Available = false
            log.Printf("Event date already passed; setting Available=false")
        } else {
            ev.Available = true
            log.Printf("Event date is upcoming; setting Available=true")
        }
    }
    
    log.Printf("=== Final event state before update ===")
    log.Printf("ID: %s", ev.ID)
    log.Printf("Title: %s", ev.Title)
    log.Printf("Category: %s", ev.Category)
    log.Printf("Date: %v", ev.Date)
    log.Printf("Quota: %d", ev.Quota)
    log.Printf("EventType: %s", ev.EventType)
    log.Printf("Price: %d", ev.Price)
    log.Printf("Available: %v", ev.Available)
    log.Printf("Rules count: %d", len(ev.Rules))
    log.Printf("Benefits count: %d", len(ev.Benefits))
    log.Printf("Poster: %s", ev.Poster)
    
    if err := s.Repo.Update(ev); err != nil {
        log.Printf("✗ Repository update error: %v", err)
        return nil, err
    }
    
    log.Printf("✓ Event updated successfully: %s", ev.ID)
    return ev, nil
}

func (s *EventService) Delete(id string) error {
    log.Printf("Deleting event: %s", id)
    return s.Repo.Delete(id)
}

func (s *EventService) GetByID(id string) (*model.Event, error) {
    log.Printf("Fetching event by ID: %s", id)
    return s.Repo.GetByID(id)
}

func (s *EventService) List(category, search string) ([]model.Event, error) {
    log.Printf("Listing events - category: %q, search: %q", category, search)
    return s.Repo.List(category, search)
}

func getString(m map[string]interface{}, k string) string {
    if v, ok := m[k]; ok {
        switch val := v.(type) {
        case string:
            return val
        case []byte:
            return string(val)
        }
    }
    return ""
}

func isDatePassed(dt *time.Time) bool {
    if dt == nil {
        return false
    }
    now := time.Now()
    y1, m1, d1 := dt.Date()
    y2, m2, d2 := now.Date()
    if y1 < y2 {
        return true
    }
    if y1 == y2 {
        if m1 < m2 {
            return true
        }
        if m1 == m2 && d1 < d2 {
            return true
        }
    }
    return false
}