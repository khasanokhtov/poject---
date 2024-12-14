package models

import "time"

type AgroOperationModel struct {
	ID                          int64                  `json:"id" gorm:"primaryKey"`
	FieldID                     int64                  `json:"field_id"`
	FieldShapeID                int64                  `json:"field_shape_id"`
	AgriWorkPlanID              *int64                 `json:"agri_work_plan_id"`
	ApplicationsType            string                 `json:"applications_type"`
	OperationNumber             string                 `json:"operation_number"`
	CustomName                  *string                `json:"custom_name"`
	WorkTypeID                  int64                  `json:"work_type_id"`
	Status                      string                 `json:"status"`
	CalcBy                      string                 `json:"calc_by"`
	PlannedArea                 float64                `json:"planned_area"`
	CompletedArea               float64                `json:"completed_area"`
	HarvestedWeight             float64                `json:"harvested_weight"`
	MarketableWeight            *float64               `json:"marketable_weight"`
	PlannedStartDate            string                 `json:"planned_start_date"`
	PlannedEndDate              string                 `json:"planned_end_date"`
	CompletedDate               string                 `json:"completed_date"`
	CompletedDatetime           time.Time              `json:"completed_datetime"`
	Season                      int                    `json:"season"`
	PlannedRowSpacing           *float64               `json:"planned_row_spacing"`
	PlannedPlantSpacing         *float64               `json:"planned_plant_spacing"`
	PlannedDepth                *float64               `json:"planned_depth"`
	PlannedSpeed                *float64               `json:"planned_speed"`
	PlannedWaterRate            float64                `json:"planned_water_rate"`
	FactWaterRate               float64                `json:"fact_water_rate"`
	CoveredArea                 float64                `json:"covered_area"`
	CoveredAreaByTrack          float64                `json:"covered_area_by_track"`
	MachineWorkArea             float64                `json:"machine_work_area"`
	AdditionalInfo              string                 `json:"additional_info"`
	Description                 string                 `json:"description"`
	CreatedAt                   time.Time              `json:"created_at"`
	UpdatedAt                   time.Time              `json:"updated_at"`
	ExternalID                  *string                `json:"external_id"`
	ActualStartDatetime         time.Time              `json:"actual_start_datetime"`
	AgroRecommendationID        *int64                 `json:"agro_recommendation_id"`
	LockedToEdit                bool                   `json:"locked_to_edit"`
	IdempotencyKey              *string                `json:"idempotency_key"`
	LockedAt                    *time.Time             `json:"locked_at"`
	ControlThreshing            float64                `json:"control_threshing"`
	Humidity                    float64                `json:"humidity"`
	ProteinContent              float64                `json:"protein_content"`
	OilContent                  float64                `json:"oil_content"`
	HistoryItemID               int64                  `json:"history_item_id"`
	NotifyResponsibleUsers      bool                   `json:"notify_responsible_users"`
	ResponsiblePersonID         *int64                 `json:"responsible_person_id"`
	GrainNature                 *float64               `json:"grain_nature"`
	HarmfulAdmixture            *float64               `json:"harmful_admixture"`
	GarbageAdmixture            *float64               `json:"garbage_admixture"`
	GrainAdmixture              *float64               `json:"grain_admixture"`
	OilAcidNumber               *float64               `json:"oil_acid_number"`
	FromWarehouseID             *int64                 `json:"from_warehouse_id"`
	ResponsibleUserIDs          []int64                `json:"responsible_user_ids" gorm:"-"`
	ClosureReasonID             *int64                 `json:"closure_reason_id"`
	AdditionalProductType       *string                `json:"additional_product_type"`
	AdditionalProductWeight     *float64               `json:"additional_product_weight"`
	ResponsibleUserMappingItems []map[string]interface{} `json:"responsible_user_mapping_items" gorm:"-"`
	XCustomFieldsData           map[string]interface{} `json:"x_custom_fields_data" gorm:"-"`
}
