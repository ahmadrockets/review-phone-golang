package controllers

import (
	"net/http"

	"final-project-sanber/models"
	"final-project-sanber/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type phoneInput struct {
	Name              string `json:"name"`
	BrandID           uint   `json:"brandID"`
	NetworkTechnology string `json:"network_technology"`
	LaunchAnnounce    string `json:"launch_announce"`
	LaunchStatus      string `json:"launch_status"`
	BodyDimensions    string `json:"body_dimensions"`
	BodyWeight        string `json:"body_weight"`
	BodySim           string `json:"body_sim"`
	DisplayType       string `json:"display_type"`
	DisplaySize       string `json:"display_size"`
	DisplayResolution string `json:"display_resolution"`
	DisplayProtection string `json:"display_protection"`
	PlatformOs        string `json:"platform_os"`
	PlatformCpu       string `json:"platform_cpu"`
	PlatformGpu       string `json:"platform_gpu"`
	MemoryCardslot    string `json:"memory_cardslot"`
	MemoryInternal    string `json:"memory_internal"`
	MaincamDual       string `json:"maincam_dual"`
	MaincamFeatures   string `json:"maincam_features"`
	MaincamVideo      string `json:"maincam_video"`
	SelfcamSingle     string `json:"selfcam_single"`
	SelfcamFeatures   string `json:"selfcam_features"`
	SelfcamVideo      string `json:"selfcam_video"`
	SoundLoudspeaker  string `json:"sound_loudspeaker"`
	SoundJack         string `json:"sound_jack"`
	CommsWlan         string `json:"comms_wlan"`
	CommsBluetooth    string `json:"comms_bluetooth"`
	CommsGps          string `json:"comms_gps"`
	CommsNfc          string `json:"comms_nfc"`
	CommsInfrared     string `json:"comms_infrared"`
	CommsRadio        string `json:"comms_radio"`
	CommsUsb          string `json:"comms_usb"`
	FeaturesSensor    string `json:"features_sensor"`
	BatteryType       string `json:"battery_type"`
	BatteryCharging   string `json:"battery_charging"`
	MiscColor         string `json:"misc_color"`
	MiscPrice         string `json:"misc_price"`
	PictureCover      string `json:"picture_cover"`
}

// GetAllPhone godoc
// @Summary Get all Phone.
// @Description Get a list of Phone.
// @Tags Phone
// @Produce json
// @Success 200 {object} []models.Phone
// @Router /phones [get]
func GetAllPhone(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var phones []models.Phone
	db.Find(&phones)

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// GetFavouritePhone godoc
// @Summary Get list Phone order by favourite, by how many hits.
// @Description Get a list of Phone.
// @Tags Phone
// @Produce json
// @Success 200 {object} []models.Phone
// @Router /phones/fav [get]
func GetFavouritePhone(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var phones []models.Phone
	db.Order("hits desc").Find(&phones)

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// GetAllPhoneByBrand godoc
// @Summary Get all Phone By Brand id.
// @Description Get a list of Phone By Brand id.
// @Tags Phone
// @Param id path string true "Brand id"
// @Produce json
// @Success 200 {object} []models.Phone
// @Router /phones/bybrand/{id} [get]
func GetAllPhoneByBrand(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var phones []models.Phone
	db.Where("brand_id = ?", c.Param("id")).Find(&phones)

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// GetAllPhoneByKeyword godoc
// @Summary Get all Phone By Name.
// @Description Get a list of Phone By Name.
// @Tags Phone
// @Param keyword path string true "Keyword"
// @Produce json
// @Success 200 {object} []models.Phone
// @Router /phones/search/{keyword} [get]
func GetAllPhoneByKeyword(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var phones []models.Phone
	db.Where("name LIKE ?", "%"+c.Param("keyword")+"%").Find(&phones)

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// GetDetailPhone godoc
// @Summary Get Detail Phone.
// @Description Get detail of Phone.
// @Tags Phone
// @Param id path string true "ID Phone"
// @Produce json
// @Success 200 {object} models.Phone
// @Router /phones/{id} [get]
func GetDetailPhone(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// update hits
	var updatedInput models.Phone
	updatedInput.Hits = phone.Hits + 1
	db.Model(&phone).Updates(updatedInput)

	// get brand
	var brand models.Brand
	db.Where("id = ?", phone.ID).First(&brand)
	phone.Brand = brand

	// get pictures
	var pictures []models.Picture
	db.Where("phone_id = ?", c.Param("id")).Find(&pictures)
	phone.Picture = pictures

	// get reviews
	var reviews []models.Review
	db.Where("phone_id = ?", c.Param("id")).Find(&reviews)
	phone.Review = reviews

	// get opinions
	var opinions []models.Opinion
	db.Where("phone_id = ?", c.Param("id")).Find(&opinions)
	phone.Opinion = opinions

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// CreatePhone godoc
// @Summary Create New Phone.
// @Description Creating a new Phone.
// @Tags Phone
// @Param Body body phoneInput true "the body to create a new Phone"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Phone
// @Router /phones [post]
func CreatePhone(c *gin.Context) {
	// Validate input
	var input phoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate url picture
	urlPictureCover := ""
	if input.PictureCover != "" {
		urlval := utils.UrlValidation(input.PictureCover)
		if urlval == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": "picture is not valid url"})
			return
		}
		urlPictureCover = input.PictureCover
	}
	// Create Phone
	phone := models.Phone{
		Name:              input.Name,
		BrandID:           input.BrandID,
		NetworkTechnology: input.NetworkTechnology,
		LaunchAnnounce:    input.LaunchAnnounce,
		LaunchStatus:      input.LaunchStatus,
		BodyDimensions:    input.BodyDimensions,
		BodyWeight:        input.BodyWeight,
		BodySim:           input.BodySim,
		DisplayType:       input.DisplayType,
		DisplaySize:       input.DisplaySize,
		DisplayResolution: input.DisplayResolution,
		DisplayProtection: input.DisplayProtection,
		PlatformOs:        input.PlatformOs,
		PlatformCpu:       input.PlatformCpu,
		PlatformGpu:       input.PlatformGpu,
		MemoryCardslot:    input.MemoryCardslot,
		MemoryInternal:    input.MemoryInternal,
		MaincamDual:       input.MaincamDual,
		MaincamFeatures:   input.MaincamFeatures,
		MaincamVideo:      input.MaincamVideo,
		SelfcamSingle:     input.SelfcamSingle,
		SelfcamFeatures:   input.SelfcamFeatures,
		SelfcamVideo:      input.SelfcamVideo,
		SoundLoudspeaker:  input.SoundLoudspeaker,
		SoundJack:         input.SoundJack,
		CommsWlan:         input.CommsWlan,
		CommsBluetooth:    input.CommsBluetooth,
		CommsGps:          input.CommsGps,
		CommsNfc:          input.CommsNfc,
		CommsInfrared:     input.CommsInfrared,
		CommsRadio:        input.CommsRadio,
		CommsUsb:          input.CommsUsb,
		FeaturesSensor:    input.FeaturesSensor,
		BatteryType:       input.BatteryType,
		BatteryCharging:   input.BatteryCharging,
		MiscColor:         input.MiscColor,
		MiscPrice:         input.MiscPrice,
		PictureCover:      urlPictureCover,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&phone)

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// UpdatePhone godoc
// @Summary Update Phone.
// @Description Update Phone by id.
// @Tags Phone
// @Produce json
// @Param id path string true "Phone id"
// @Param Body body phoneInput true "the body to update age rating category"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Phone
// @Router /phones/{id} [patch]
func UpdatePhone(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input phoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate url picture
	urlPictureCover := ""
	if input.PictureCover != "" {
		urlval := utils.UrlValidation(input.PictureCover)
		if urlval == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": "picture is not valid url"})
			return
		}
		urlPictureCover = input.PictureCover
	}

	var updatedInput models.Phone
	updatedInput.Name = input.Name
	updatedInput.BrandID = input.BrandID
	updatedInput.NetworkTechnology = input.NetworkTechnology
	updatedInput.LaunchAnnounce = input.LaunchAnnounce
	updatedInput.LaunchStatus = input.LaunchStatus
	updatedInput.BodyDimensions = input.BodyDimensions
	updatedInput.BodyWeight = input.BodyWeight
	updatedInput.BodySim = input.BodySim
	updatedInput.DisplayType = input.DisplayType
	updatedInput.DisplaySize = input.DisplaySize
	updatedInput.DisplayResolution = input.DisplayResolution
	updatedInput.DisplayProtection = input.DisplayProtection
	updatedInput.PlatformOs = input.PlatformOs
	updatedInput.PlatformCpu = input.PlatformCpu
	updatedInput.PlatformGpu = input.PlatformGpu
	updatedInput.MemoryCardslot = input.MemoryCardslot
	updatedInput.MemoryInternal = input.MemoryInternal
	updatedInput.MaincamDual = input.MaincamDual
	updatedInput.MaincamFeatures = input.MaincamFeatures
	updatedInput.MaincamVideo = input.MaincamVideo
	updatedInput.SelfcamSingle = input.SelfcamSingle
	updatedInput.SelfcamFeatures = input.SelfcamFeatures
	updatedInput.SelfcamVideo = input.SelfcamVideo
	updatedInput.SoundLoudspeaker = input.SoundLoudspeaker
	updatedInput.SoundJack = input.SoundJack
	updatedInput.CommsWlan = input.CommsWlan
	updatedInput.CommsBluetooth = input.CommsBluetooth
	updatedInput.CommsGps = input.CommsGps
	updatedInput.CommsNfc = input.CommsNfc
	updatedInput.CommsInfrared = input.CommsInfrared
	updatedInput.CommsRadio = input.CommsRadio
	updatedInput.CommsUsb = input.CommsUsb
	updatedInput.FeaturesSensor = input.FeaturesSensor
	updatedInput.BatteryType = input.BatteryType
	updatedInput.BatteryCharging = input.BatteryCharging
	updatedInput.MiscColor = input.MiscColor
	updatedInput.MiscPrice = input.MiscPrice
	updatedInput.PictureCover = urlPictureCover

	db.Model(&phone).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": phone})
}

// DeletePhone godoc
// @Summary Delete one Phone.
// @Description Delete a Phone by id.
// @Tags Phone
// @Produce json
// @Param id path string true "Phone id"
// @Param Authorization header string true "Authorization. How to input in swagger: 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /phones/{id} [delete]
func DeletePhone(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var phone models.Phone
	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&phone)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
