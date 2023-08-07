package models

// BrandAndExternalIds represents a struct with BrandID and ExternalID fields.
type BrandAndExternalIds struct {
	BrandID    int64
	ExternalID string
}

// NewBrandAndExternalIds creates a new instance of BrandAndExternalIds.
func NewBrandAndExternalIds(brandId int64, externalId string) *BrandAndExternalIds {
	return &BrandAndExternalIds{
		BrandID:    brandId,
		ExternalID: externalId,
	}
}

// BrandAndExternalIdsBuilder is a builder for BrandAndExternalIds.
type BrandAndExternalIdsBuilder struct {
	BrandId    int64
	ExternalId string
}

// WithBrandId sets the BrandId field for the builder.
func (b *BrandAndExternalIdsBuilder) WithBrandId(brandId int64) *BrandAndExternalIdsBuilder {
	b.BrandId = brandId
	return b
}

// WithExternalId sets the ExternalId field for the builder.
func (b *BrandAndExternalIdsBuilder) WithExternalId(externalId string) *BrandAndExternalIdsBuilder {
	b.ExternalId = externalId
	return b
}

// Build creates a new instance of BrandAndExternalIds using the values from the builder.
func (b *BrandAndExternalIdsBuilder) Build() *BrandAndExternalIds {
	return &BrandAndExternalIds{
		BrandID:    b.BrandId,
		ExternalID: b.ExternalId,
	}
}
