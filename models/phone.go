package models

import "time"

type (
	// Phone
	Phone struct {
		ID                uint      `json:"id" gorm:"primary_key"`
		Name              string    `gorm:"not null;unique" json:"name"`
		BrandID           uint      `json:"brandID"`
		NetworkTechnology string    `json:"network_technology"`
		LaunchAnnounce    string    `json:"launch_announce"`
		LaunchStatus      string    `json:"launch_status"`
		BodyDimensions    string    `json:"body_dimensions"`
		BodyWeight        string    `json:"body_weight"`
		BodySim           string    `json:"body_sim"`
		DisplayType       string    `json:"display_type"`
		DisplaySize       string    `json:"display_size"`
		DisplayResolution string    `json:"display_resolution"`
		DisplayProtection string    `json:"display_protection"`
		PlatformOs        string    `json:"platform_os"`
		PlatformCpu       string    `json:"platform_cpu"`
		PlatformGpu       string    `json:"platform_gpu"`
		MemoryCardslot    string    `json:"memory_cardslot"`
		MemoryInternal    string    `json:"memory_internal"`
		MaincamDual       string    `json:"maincam_dual"`
		MaincamFeatures   string    `json:"maincam_features"`
		MaincamVideo      string    `json:"maincam_video"`
		SelfcamSingle     string    `json:"selfcam_single"`
		SelfcamFeatures   string    `json:"selfcam_features"`
		SelfcamVideo      string    `json:"selfcam_video"`
		SoundLoudspeaker  string    `json:"sound_loudspeaker"`
		SoundJack         string    `json:"sound_jack"`
		CommsWlan         string    `json:"comms_wlan"`
		CommsBluetooth    string    `json:"comms_bluetooth"`
		CommsGps          string    `json:"comms_gps"`
		CommsNfc          string    `json:"comms_nfc"`
		CommsInfrared     string    `json:"comms_infrared"`
		CommsRadio        string    `json:"comms_radio"`
		CommsUsb          string    `json:"comms_usb"`
		FeaturesSensor    string    `json:"features_sensor"`
		BatteryType       string    `json:"battery_type"`
		BatteryCharging   string    `json:"battery_charging"`
		MiscColor         string    `json:"misc_color"`
		MiscPrice         string    `json:"misc_price"`
		PictureCover      string    `json:"picture_cover"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Hits              int       `json:"hits"`
		Brand             Brand     `json:"brand"`
		Review            []Review  `json:"reviews"`
		Opinion           []Opinion `json:"opinions"`
		Picture           []Picture `json:"pictures"`
	}
)
