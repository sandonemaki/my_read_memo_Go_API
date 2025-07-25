// Code generated by BobGen psql v0.38.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import "context"

type Factory struct {
	baseAuthorMods          AuthorModSlice
	baseGooseDBVersionMods  GooseDBVersionModSlice
	baseKindleHighlightMods KindleHighlightModSlice
	baseMasterBookMods      MasterBookModSlice
	baseOcrTextMods         OcrTextModSlice
	basePublisherMods       PublisherModSlice
	baseRandokuImageMods    RandokuImageModSlice
	baseRandokuMemoMods     RandokuMemoModSlice
	baseReadingHistoryMods  ReadingHistoryModSlice
	baseSeidokuMemoMods     SeidokuMemoModSlice
	baseUserBookLogMods     UserBookLogModSlice
	baseUserMods            UserModSlice
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) NewAuthor(ctx context.Context, mods ...AuthorMod) *AuthorTemplate {
	o := &AuthorTemplate{f: f}

	if f != nil {
		f.baseAuthorMods.Apply(ctx, o)
	}

	AuthorModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewGooseDBVersion(ctx context.Context, mods ...GooseDBVersionMod) *GooseDBVersionTemplate {
	o := &GooseDBVersionTemplate{f: f}

	if f != nil {
		f.baseGooseDBVersionMods.Apply(ctx, o)
	}

	GooseDBVersionModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewKindleHighlight(ctx context.Context, mods ...KindleHighlightMod) *KindleHighlightTemplate {
	o := &KindleHighlightTemplate{f: f}

	if f != nil {
		f.baseKindleHighlightMods.Apply(ctx, o)
	}

	KindleHighlightModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewMasterBook(ctx context.Context, mods ...MasterBookMod) *MasterBookTemplate {
	o := &MasterBookTemplate{f: f}

	if f != nil {
		f.baseMasterBookMods.Apply(ctx, o)
	}

	MasterBookModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewOcrText(ctx context.Context, mods ...OcrTextMod) *OcrTextTemplate {
	o := &OcrTextTemplate{f: f}

	if f != nil {
		f.baseOcrTextMods.Apply(ctx, o)
	}

	OcrTextModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewPublisher(ctx context.Context, mods ...PublisherMod) *PublisherTemplate {
	o := &PublisherTemplate{f: f}

	if f != nil {
		f.basePublisherMods.Apply(ctx, o)
	}

	PublisherModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewRandokuImage(ctx context.Context, mods ...RandokuImageMod) *RandokuImageTemplate {
	o := &RandokuImageTemplate{f: f}

	if f != nil {
		f.baseRandokuImageMods.Apply(ctx, o)
	}

	RandokuImageModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewRandokuMemo(ctx context.Context, mods ...RandokuMemoMod) *RandokuMemoTemplate {
	o := &RandokuMemoTemplate{f: f}

	if f != nil {
		f.baseRandokuMemoMods.Apply(ctx, o)
	}

	RandokuMemoModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewReadingHistory(ctx context.Context, mods ...ReadingHistoryMod) *ReadingHistoryTemplate {
	o := &ReadingHistoryTemplate{f: f}

	if f != nil {
		f.baseReadingHistoryMods.Apply(ctx, o)
	}

	ReadingHistoryModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewSeidokuMemo(ctx context.Context, mods ...SeidokuMemoMod) *SeidokuMemoTemplate {
	o := &SeidokuMemoTemplate{f: f}

	if f != nil {
		f.baseSeidokuMemoMods.Apply(ctx, o)
	}

	SeidokuMemoModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewUserBookLog(ctx context.Context, mods ...UserBookLogMod) *UserBookLogTemplate {
	o := &UserBookLogTemplate{f: f}

	if f != nil {
		f.baseUserBookLogMods.Apply(ctx, o)
	}

	UserBookLogModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) NewUser(ctx context.Context, mods ...UserMod) *UserTemplate {
	o := &UserTemplate{f: f}

	if f != nil {
		f.baseUserMods.Apply(ctx, o)
	}

	UserModSlice(mods).Apply(ctx, o)

	return o
}

func (f *Factory) ClearBaseAuthorMods() {
	f.baseAuthorMods = nil
}

func (f *Factory) AddBaseAuthorMod(mods ...AuthorMod) {
	f.baseAuthorMods = append(f.baseAuthorMods, mods...)
}

func (f *Factory) ClearBaseGooseDBVersionMods() {
	f.baseGooseDBVersionMods = nil
}

func (f *Factory) AddBaseGooseDBVersionMod(mods ...GooseDBVersionMod) {
	f.baseGooseDBVersionMods = append(f.baseGooseDBVersionMods, mods...)
}

func (f *Factory) ClearBaseKindleHighlightMods() {
	f.baseKindleHighlightMods = nil
}

func (f *Factory) AddBaseKindleHighlightMod(mods ...KindleHighlightMod) {
	f.baseKindleHighlightMods = append(f.baseKindleHighlightMods, mods...)
}

func (f *Factory) ClearBaseMasterBookMods() {
	f.baseMasterBookMods = nil
}

func (f *Factory) AddBaseMasterBookMod(mods ...MasterBookMod) {
	f.baseMasterBookMods = append(f.baseMasterBookMods, mods...)
}

func (f *Factory) ClearBaseOcrTextMods() {
	f.baseOcrTextMods = nil
}

func (f *Factory) AddBaseOcrTextMod(mods ...OcrTextMod) {
	f.baseOcrTextMods = append(f.baseOcrTextMods, mods...)
}

func (f *Factory) ClearBasePublisherMods() {
	f.basePublisherMods = nil
}

func (f *Factory) AddBasePublisherMod(mods ...PublisherMod) {
	f.basePublisherMods = append(f.basePublisherMods, mods...)
}

func (f *Factory) ClearBaseRandokuImageMods() {
	f.baseRandokuImageMods = nil
}

func (f *Factory) AddBaseRandokuImageMod(mods ...RandokuImageMod) {
	f.baseRandokuImageMods = append(f.baseRandokuImageMods, mods...)
}

func (f *Factory) ClearBaseRandokuMemoMods() {
	f.baseRandokuMemoMods = nil
}

func (f *Factory) AddBaseRandokuMemoMod(mods ...RandokuMemoMod) {
	f.baseRandokuMemoMods = append(f.baseRandokuMemoMods, mods...)
}

func (f *Factory) ClearBaseReadingHistoryMods() {
	f.baseReadingHistoryMods = nil
}

func (f *Factory) AddBaseReadingHistoryMod(mods ...ReadingHistoryMod) {
	f.baseReadingHistoryMods = append(f.baseReadingHistoryMods, mods...)
}

func (f *Factory) ClearBaseSeidokuMemoMods() {
	f.baseSeidokuMemoMods = nil
}

func (f *Factory) AddBaseSeidokuMemoMod(mods ...SeidokuMemoMod) {
	f.baseSeidokuMemoMods = append(f.baseSeidokuMemoMods, mods...)
}

func (f *Factory) ClearBaseUserBookLogMods() {
	f.baseUserBookLogMods = nil
}

func (f *Factory) AddBaseUserBookLogMod(mods ...UserBookLogMod) {
	f.baseUserBookLogMods = append(f.baseUserBookLogMods, mods...)
}

func (f *Factory) ClearBaseUserMods() {
	f.baseUserMods = nil
}

func (f *Factory) AddBaseUserMod(mods ...UserMod) {
	f.baseUserMods = append(f.baseUserMods, mods...)
}
