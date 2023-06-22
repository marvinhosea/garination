// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: model.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Dealership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealershipId string                 `protobuf:"bytes,1,opt,name=dealership_id,json=dealershipId,proto3" json:"dealership_id,omitempty"`
	OwnerId      string                 `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Name         string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName  string                 `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Address      string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	City         string                 `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	State        string                 `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Zip          string                 `protobuf:"bytes,8,opt,name=zip,proto3" json:"zip,omitempty"`
	Phone        string                 `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	Email        string                 `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Website      string                 `protobuf:"bytes,11,opt,name=website,proto3" json:"website,omitempty"`
	FacebookUrl  string                 `protobuf:"bytes,12,opt,name=facebook_url,json=facebookUrl,proto3" json:"facebook_url,omitempty"`
	TwitterUrl   string                 `protobuf:"bytes,13,opt,name=twitter_url,json=twitterUrl,proto3" json:"twitter_url,omitempty"`
	InstagramUrl string                 `protobuf:"bytes,14,opt,name=instagram_url,json=instagramUrl,proto3" json:"instagram_url,omitempty"`
	LinkedinUrl  string                 `protobuf:"bytes,15,opt,name=linkedin_url,json=linkedinUrl,proto3" json:"linkedin_url,omitempty"`
	LogoUrl      string                 `protobuf:"bytes,16,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	CoverUrl     string                 `protobuf:"bytes,17,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Description  string                 `protobuf:"bytes,18,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Dealership) Reset() {
	*x = Dealership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dealership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dealership) ProtoMessage() {}

func (x *Dealership) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dealership.ProtoReflect.Descriptor instead.
func (*Dealership) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *Dealership) GetDealershipId() string {
	if x != nil {
		return x.DealershipId
	}
	return ""
}

func (x *Dealership) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Dealership) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dealership) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Dealership) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Dealership) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Dealership) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Dealership) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *Dealership) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Dealership) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Dealership) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Dealership) GetFacebookUrl() string {
	if x != nil {
		return x.FacebookUrl
	}
	return ""
}

func (x *Dealership) GetTwitterUrl() string {
	if x != nil {
		return x.TwitterUrl
	}
	return ""
}

func (x *Dealership) GetInstagramUrl() string {
	if x != nil {
		return x.InstagramUrl
	}
	return ""
}

func (x *Dealership) GetLinkedinUrl() string {
	if x != nil {
		return x.LinkedinUrl
	}
	return ""
}

func (x *Dealership) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Dealership) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Dealership) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dealership) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Dealership) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserMetum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMetaId   string `protobuf:"bytes,1,opt,name=user_meta_id,json=userMetaId,proto3" json:"user_meta_id,omitempty"`
	UserId       string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FacebookUrl  string `protobuf:"bytes,3,opt,name=facebook_url,json=facebookUrl,proto3" json:"facebook_url,omitempty"`
	TwitterUrl   string `protobuf:"bytes,4,opt,name=twitter_url,json=twitterUrl,proto3" json:"twitter_url,omitempty"`
	InstagramUrl string `protobuf:"bytes,5,opt,name=instagram_url,json=instagramUrl,proto3" json:"instagram_url,omitempty"`
	LinkedinUrl  string `protobuf:"bytes,6,opt,name=linkedin_url,json=linkedinUrl,proto3" json:"linkedin_url,omitempty"`
	WebsiteUrl   string `protobuf:"bytes,7,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	DealershipId string `protobuf:"bytes,8,opt,name=dealership_id,json=dealershipId,proto3" json:"dealership_id,omitempty"`
}

func (x *UserMetum) Reset() {
	*x = UserMetum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMetum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMetum) ProtoMessage() {}

func (x *UserMetum) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMetum.ProtoReflect.Descriptor instead.
func (*UserMetum) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *UserMetum) GetUserMetaId() string {
	if x != nil {
		return x.UserMetaId
	}
	return ""
}

func (x *UserMetum) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserMetum) GetFacebookUrl() string {
	if x != nil {
		return x.FacebookUrl
	}
	return ""
}

func (x *UserMetum) GetTwitterUrl() string {
	if x != nil {
		return x.TwitterUrl
	}
	return ""
}

func (x *UserMetum) GetInstagramUrl() string {
	if x != nil {
		return x.InstagramUrl
	}
	return ""
}

func (x *UserMetum) GetLinkedinUrl() string {
	if x != nil {
		return x.LinkedinUrl
	}
	return ""
}

func (x *UserMetum) GetWebsiteUrl() string {
	if x != nil {
		return x.WebsiteUrl
	}
	return ""
}

func (x *UserMetum) GetDealershipId() string {
	if x != nil {
		return x.DealershipId
	}
	return ""
}

type CarBrand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId   string                 `protobuf:"bytes,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LogoUrl   string                 `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CarBrand) Reset() {
	*x = CarBrand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarBrand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarBrand) ProtoMessage() {}

func (x *CarBrand) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarBrand.ProtoReflect.Descriptor instead.
func (*CarBrand) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *CarBrand) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *CarBrand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CarBrand) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *CarBrand) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CarBrand) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CarImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarImageId string                 `protobuf:"bytes,1,opt,name=car_image_id,json=carImageId,proto3" json:"car_image_id,omitempty"`
	CarId      string                 `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Url        string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CarImage) Reset() {
	*x = CarImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarImage) ProtoMessage() {}

func (x *CarImage) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarImage.ProtoReflect.Descriptor instead.
func (*CarImage) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

func (x *CarImage) GetCarImageId() string {
	if x != nil {
		return x.CarImageId
	}
	return ""
}

func (x *CarImage) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *CarImage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CarImage) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CarImage) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId          string                 `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	BrandId        string                 `protobuf:"bytes,2,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	CategoryId     string                 `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Model          string                 `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Year           int32                  `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Price          float64                `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Mileage        int32                  `protobuf:"varint,7,opt,name=mileage,proto3" json:"mileage,omitempty"`
	Color          string                 `protobuf:"bytes,8,opt,name=color,proto3" json:"color,omitempty"`
	Transmission   string                 `protobuf:"bytes,9,opt,name=transmission,proto3" json:"transmission,omitempty"`
	FuelType       string                 `protobuf:"bytes,10,opt,name=fuel_type,json=fuelType,proto3" json:"fuel_type,omitempty"`
	EngineCapacity string                 `protobuf:"bytes,11,opt,name=engine_capacity,json=engineCapacity,proto3" json:"engine_capacity,omitempty"`
	Description    string                 `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	DealershipId   string                 `protobuf:"bytes,13,opt,name=dealership_id,json=dealershipId,proto3" json:"dealership_id,omitempty"`
	DealerId       string                 `protobuf:"bytes,14,opt,name=dealer_id,json=dealerId,proto3" json:"dealer_id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{4}
}

func (x *Car) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *Car) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *Car) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Car) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Car) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Car) GetMileage() int32 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

func (x *Car) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Car) GetTransmission() string {
	if x != nil {
		return x.Transmission
	}
	return ""
}

func (x *Car) GetFuelType() string {
	if x != nil {
		return x.FuelType
	}
	return ""
}

func (x *Car) GetEngineCapacity() string {
	if x != nil {
		return x.EngineCapacity
	}
	return ""
}

func (x *Car) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Car) GetDealershipId() string {
	if x != nil {
		return x.DealershipId
	}
	return ""
}

func (x *Car) GetDealerId() string {
	if x != nil {
		return x.DealerId
	}
	return ""
}

func (x *Car) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Car) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SingleCar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car    *Car               `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	Images []*CarImage        `protobuf:"bytes,2,rep,name=images,proto3" json:"images,omitempty"`
	Extras []*CarExtraFeature `protobuf:"bytes,3,rep,name=extras,proto3" json:"extras,omitempty"`
	Brand  *CarBrand          `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *SingleCar) Reset() {
	*x = SingleCar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleCar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleCar) ProtoMessage() {}

func (x *SingleCar) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleCar.ProtoReflect.Descriptor instead.
func (*SingleCar) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{5}
}

func (x *SingleCar) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *SingleCar) GetImages() []*CarImage {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *SingleCar) GetExtras() []*CarExtraFeature {
	if x != nil {
		return x.Extras
	}
	return nil
}

func (x *SingleCar) GetBrand() *CarBrand {
	if x != nil {
		return x.Brand
	}
	return nil
}

type CarExtraFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarExtraId string                 `protobuf:"bytes,1,opt,name=car_extra_id,json=carExtraId,proto3" json:"car_extra_id,omitempty"`
	CarId      string                 `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Value      string                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CarExtraFeature) Reset() {
	*x = CarExtraFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarExtraFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarExtraFeature) ProtoMessage() {}

func (x *CarExtraFeature) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarExtraFeature.ProtoReflect.Descriptor instead.
func (*CarExtraFeature) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{6}
}

func (x *CarExtraFeature) GetCarExtraId() string {
	if x != nil {
		return x.CarExtraId
	}
	return ""
}

func (x *CarExtraFeature) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *CarExtraFeature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CarExtraFeature) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CarExtraFeature) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CarExtraFeature) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb,
	0x04, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x98, 0x02, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61, 0x63,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x77, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xcb, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x8c, 0x04, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61,
	0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x12,
	0x16, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x43,
	0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x72, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x61, 0x72,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x06, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0xea, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72,
	0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x61, 0x72, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x62, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_model_proto_goTypes = []interface{}{
	(*Dealership)(nil),            // 0: Dealership
	(*UserMetum)(nil),             // 1: UserMetum
	(*CarBrand)(nil),              // 2: CarBrand
	(*CarImage)(nil),              // 3: CarImage
	(*Car)(nil),                   // 4: Car
	(*SingleCar)(nil),             // 5: SingleCar
	(*CarExtraFeature)(nil),       // 6: CarExtraFeature
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_model_proto_depIdxs = []int32{
	7,  // 0: Dealership.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: Dealership.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 2: CarBrand.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: CarBrand.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 4: CarImage.created_at:type_name -> google.protobuf.Timestamp
	7,  // 5: CarImage.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 6: Car.created_at:type_name -> google.protobuf.Timestamp
	7,  // 7: Car.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 8: SingleCar.car:type_name -> Car
	3,  // 9: SingleCar.images:type_name -> CarImage
	6,  // 10: SingleCar.extras:type_name -> CarExtraFeature
	2,  // 11: SingleCar.brand:type_name -> CarBrand
	7,  // 12: CarExtraFeature.created_at:type_name -> google.protobuf.Timestamp
	7,  // 13: CarExtraFeature.updated_at:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dealership); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMetum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarBrand); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarImage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleCar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarExtraFeature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
