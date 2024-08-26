package Logic

import (
	"errors"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/DTO"

	"gorm.io/gorm"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(db *gorm.DB, data Const.Paging) ([]*DTO.KolDTO, int64, error) {
	if data.PageIndex <= 0 || data.PageSize <= 0 {
		return nil, 0, errors.New("invalid pageIndex or pageSize")
	}

	// Calculate offset for pagination
	offset := (data.PageIndex - 1) * data.PageSize

	var kols []*DTO.KolDTO
	var totalCount int64
	var queries string = "SELECT kolid, userprofileid, language, education, expectedsalary, expectedsalaryenable, channelsettingtypeid, idfronturl, idbackurl, portraiturl, rewardid, paymentmethodid, testimonialsid, verificationstatus, enabled, activedate, active, createdby, createddate, modifiedby, modifieddate, isremove, isonboarding, code, portraitrighturl, portraitlefturl, livenessstatus FROM kol_dtos LIMIT ? OFFSET ?"

	rows, err := db.Raw(queries, data.PageSize, offset).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var kol DTO.KolDTO
		err := rows.Scan(&kol.KolID, &kol.UserProfileID, &kol.Language, &kol.Education, &kol.ExpectedSalary, &kol.ExpectedSalaryEnable, &kol.ChannelSettingTypeID, &kol.IDFrontURL, &kol.IDBackURL, &kol.PortraitURL, &kol.RewardID, &kol.PaymentMethodID, &kol.TestimonialsID, &kol.VerificationStatus, &kol.Enabled, &kol.ActiveDate, &kol.Active, &kol.CreatedBy, &kol.CreatedDate, &kol.ModifiedBy, &kol.ModifiedDate, &kol.IsRemove, &kol.IsOnBoarding, &kol.Code, &kol.PortraitRightURL, &kol.PortraitLeftURL, &kol.LivenessStatus)
		if err != nil {
			return nil, 0, err
		}
		kols = append(kols, &kol)
	}
	totalCount = int64(len(kols))
	return kols, totalCount, nil
}
