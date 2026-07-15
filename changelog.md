## v15.0.0 - 2026-07-15
### Breaking Changes
* **`EarnedRewardApprovedData.Attributes`** — type changed from `*RewardNotificationAttributes` to `*EarnedRewardNotificationAttributes`; update all type references and any variables typed as `*RewardNotificationAttributes` to use `*EarnedRewardNotificationAttributes` instead.
* **`EarnedRewardApprovedData.GetAttributes()`** and **`SetAttributes()`** — return/parameter type updated to `*EarnedRewardNotificationAttributes`; callers must update their type assertions and assignments.
### Added
* **`EarnedRewardNotificationAttributes`** — new struct replacing `RewardNotificationAttributes` for approved-reward notifications, carrying message, merchant name, attribution URL, transaction details, category, user reward, assets, and purchase channels.
* **`UserReward`** — new struct representing the commission type (`CommissionType`) and value (`float64`) for an offer reward, surfaced on both `EarnedRewardNotificationAttributes` and `EarnedRewardSettledAttributes`.
* **`EarnedRewardSettledAttributes`** — four new optional fields added: `CategoryName`, `UserReward`, `Assets`, and `PurchaseChannel`, with corresponding getters and setters.

## v14.0.0 - 2026-07-15
### Breaking Changes
* **`NotificationDataUnionVisitor`** — new required method `VisitEarnedRewardRejected(*EarnedRewardRejectedData) error` added to the interface; all existing implementations must add this method to continue compiling.
### Added
* **`NotificationTypeEarnedRewardRejected`** — new `NotificationType` enum constant (`"earnedRewardRejected"`) recognized by `NewNotificationTypeFromString`.
* **`EarnedRewardRejectedAttributes`** — new struct carrying the reason code, display message, transaction ID, amount, and optional timestamp for a rejected reward notification.
* **`EarnedRewardRejectedData`** — new struct wrapping `EarnedRewardRejectedAttributes` and `RejectedTransactionRelationships`, surfaced as the `EarnedRewardRejected` field on `NotificationDataUnion`.
* **`RejectedTransactionRelationships`** — new struct linking the user and transaction relationship references for a rejected-reward notification.

## v13.4.0 - 2026-07-10
### Added
* **`ContentStrategySortOffersNearYou`** — new `ContentStrategySort` enum value (`"OFFERS_NEAR_YOU"`) representing a location-based offers sort; recognized by `NewContentStrategySortFromString` for round-trip serialization.

## v13.3.0 - 2026-07-02
### Added
* **`ProgressBarSegmentProgress`** — new type representing the fill state of a single progress bar segment node, with `Completed` and `Total` integer fields and full JSON serialization support.
* **`ProgressBarSegments.Progress`** — new field (`[]*ProgressBarSegmentProgress`) exposing per-segment fill state, index-aligned with the segment nodes array; accessible via `GetProgress` and `SetProgress`.

## v13.2.1 - 2026-07-02
* chore: update User-Agent header placeholder and clarify FilterSearch doc
* Update the User-Agent header value to the Fern version placeholder and
* improve the godoc comment for the FilterSearch field on
* GetOffersByUserRequest to more accurately describe its matching behavior.
* Key changes:
* Replace versioned User-Agent string with "0.0.0-fern-placeholder" in request_option.go
* Clarify FilterSearch comment: now documents substring matching on offer name or category name, not just merchant name
* 🌿 Generated with Fern

## v13.2.0 - 2026-06-29
### Added
* **`OfferMediumPush`** — new `OfferMedium` enum value (`"PUSH"`) representing the push channel; recognized by `NewOfferMediumFromString` for round-trip serialization.

## v13.1.0 - 2026-06-23
### Added
* **`BatchesMeta`** — new type representing metadata about a batch placement, with `PlacementName` field and getter/setter methods.
* **`BatchesResponseObject.Meta`** — new optional field of type `BatchesMeta` providing placement metadata in batch responses.
* **`OffersMeta.PlacementName`** — new optional field exposing the display name of the placement resolved server-side; populated on the Get Placement Content endpoint.

## v13.0.0 - 2026-06-23
### Breaking Changes
* **`Client.PlacementOffers`** and **`Client.PlacementBatches`** methods (and their `RawClient` counterparts) have been removed; migrate to the unified **`Client.PlacementContent`** endpoint, which resolves the placement type server-side and returns either offers or batches automatically.
* **`GetOffersByPlacementRequest`** and **`GetBatchesByPlacementRequest`** types have been removed; use **`GetPlacementContentRequest`** with `PlacementContent` instead.
* **`PlacementContentData`** union type and **`PlacementContentDataVisitor`** interface have been removed and replaced by **`PlacementContentResponse`** (a union of `OffersResponseObject` / `BatchesResponseObject`) and **`PlacementContentResponseVisitor`** (with `VisitOffersResponseObject` / `VisitBatchesResponseObject` methods); update all type assertions, visitor implementations, and `Accept` call sites accordingly.

## v12.2.0 - 2026-06-22
### Added
* **`Client.PlacementContent`** — new method that retrieves placement content for a given organization, user, and placement ID; the server resolves the placement type automatically, returning `standardOffer` resources for standard placements and `placementBatch` slot resources for batch-activation or group placements.
* **`GetPlacementContentRequest`** — new request type with optional `Include` and `SupportedComponents` fields (and `SetInclude`/`SetSupportedComponents` setters) used to parameterize the `PlacementContent` call.
* **`PlacementContentData`** — new union type that wraps either an `OfferDataUnion` or a `PlacementBatchData` resource, with `GetOfferDataUnion()`, `GetPlacementBatchData()`, and `Accept(PlacementContentDataVisitor)` for exhaustive dispatch.
* **`PlacementContentResponse`** — new JSON:API response envelope with `Data`, `Links`, `Included`, and `Meta` fields, mirroring the shape of the existing offers and batches responses.

## v12.1.0 - 2026-06-17
### Added
* **`TransactionsAttributes.AccountId`** — new optional `*string` field representing an account identifier associated with a transaction, accessible via `GetAccountId()` and settable via `SetAccountId()`.

## v12.0.0 - 2026-06-11
* feat!: rename MainPage/BatchActivationRelationships types and add email/group placement support
* This release introduces two new placement types (email and group), renames
* the `placementMainPage` discriminant to `placement` across all union types,
* and replaces several `MainPage`-prefixed types with more generic equivalents.
* Existing implementations of `CreatePlacementDataUnionVisitor`,
* `UpdatePlacementDataUnionVisitor`, `PlacementFormatUnionVisitor`, and
* `IncludedResourceVisitor` must be updated to satisfy the revised interfaces.
* Key changes:
* Removed `BatchActivationPlacementRelationships`; replaced by new `SlottedPlacementRelationships` (used by both batch-activation and group placements)
* Renamed `MainPagePlacementAttributes` → `PlacementAttributes`, `MainPagePlacementData` → `PlacementData`, `CreateMainPageAttributes` → `CreateStandardAttributes`, `CreateMainPagePlacementData` → `CreateStandardPlacementData`, `UpdateMainPageAttributes` → `UpdateStandardAttributes`, `UpdateMainPagePlacementData` → `UpdateStandardPlacementData`
* Renamed visitor methods `VisitPlacementMainPage` → `VisitPlacement` and union fields `PlacementMainPage` → `Placement` across `CreatePlacementDataUnion`, `UpdatePlacementDataUnion`, `PlacementFormatUnion`, and `IncludedResource`
* Added new `placementEmail` and `placementGroup` placement types with full create/update/read type families (`CreateEmailAttributes`, `CreateEmailPlacementData`, `EmailPlacementAttributes`, `EmailPlacementData`, `UpdateEmailAttributes`, `UpdateEmailPlacementData`, `CreateGroupAttributes`, `CreateGroupPlacementData`, `GroupPlacementAttributes`, `GroupPlacementData`, `UpdateGroupAttributes`, `UpdateGroupPlacementData`)
* Added `PlacementTypeFilterPlacementEmail` and `PlacementTypeFilterPlacementGroup` constants; renamed `PlacementTypeFilterPlacementMainPage` → `PlacementTypeFilterPlacement`
* 🌿 Generated with Fern

## v11.0.0 - 2026-06-10
### Breaking Changes
* **`NotificationDataUnionVisitor`** — two new required methods `VisitPushNotificationPlacementFile(*PushNotificationPlacementFileData) error` and `VisitEmailNotificationPlacementFile(*EmailNotificationPlacementFileData) error` have been added to the interface; all existing implementations must add these methods to continue compiling.
### Added
* **`NotificationTypePushNotificationPlacementFile`** and **`NotificationTypeEmailNotificationPlacementFile`** — new `NotificationType` constants for placement file webhook notifications.
* **`PushNotificationPlacementFileData`**, **`PushNotificationPlacementFileAttributes`**, and **`PushNotificationPlacementFileRelationships`** — new types representing a push-notification placement file notification payload, including presigned download URL, cadence, and available slot count.
* **`EmailNotificationPlacementFileData`**, **`EmailNotificationPlacementFileAttributes`**, and **`EmailNotificationPlacementFileRelationships`** — new types representing an email placement file notification payload, including organization ID, presigned download URL, cadence, and available slot count.
* **`NotificationDataUnion.PushNotificationPlacementFile`** and **`NotificationDataUnion.EmailNotificationPlacementFile`** — new union fields with corresponding `GetPushNotificationPlacementFile()` and `GetEmailNotificationPlacementFile()` accessors.

## v10.1.0 - 2026-06-10
### Added
* **`UpdateUserRequestAttributes.HistoricalTransactionsSent`** — new optional `*bool` field that confirms historical transactions have been sent for a user; once set to `true` it cannot be reverted, accessible via `GetHistoricalTransactionsSent` and `SetHistoricalTransactionsSent`.

## v10.0.0 - 2026-06-01
### Breaking Changes
* **`PlacementBatchAttributes.ShortDescription`** — field and its `GetShortDescription`/`SetShortDescription` methods are removed; migrate by reading `shortDescription` from the `Components` (`*OfferComponents`) field instead.
* **`PlacementBatchAttributes.LongDescription`** — field and its `GetLongDescription`/`SetLongDescription` methods are removed; migrate by reading `longDescription` from the `Components` (`*OfferComponents`) field instead.

## v9.1.0 - 2026-06-01
### Added
* **`PlacementBatchAttributes.ShortDescription`** — new string field with a brief, human-readable label describing how long a slot stays activated after a user taps activate (e.g. `"Activated for 24 hours"`), accessible via `GetShortDescription` and `SetShortDescription`.
* **`PlacementBatchAttributes.LongDescription`** — new string field with a fuller description of the slot's activation behavior, clarifying which offers are activated and for how long, accessible via `GetLongDescription` and `SetLongDescription`.

## v9.0.0 - 2026-06-01
### Breaking Changes
* **`BatchSlotData`** struct is removed; migrate by using the new `PlacementBatchData` (which exposes `Id` and `Attributes`) and `PlacementBatchAttributes` (which carries `Name`, `IsActive`, `LastActivatedAt`, `ExpiresAt`, `Components`, `Assets`, and `Offers`) instead.
* **`BatchesResponseObject.Data`** type changed from `[]*BatchSlotData` to `[]*PlacementBatchData`; update all call sites that read or write this field, including `GetData()` and `SetData()`.
### Added
* **`PlacementBatchData`** — new JSON:API resource struct representing a batch-activation slot, with `Id` (stable slot identifier), a `type` literal of `"placementBatch"`, and an `Attributes` field pointing to `PlacementBatchAttributes`.
* **`PlacementBatchAttributes`** — new struct holding slot-level detail: `Name`, `IsActive`, `LastActivatedAt`, `ExpiresAt`, `Components`, `Assets`, and `Offers`, with corresponding getters and setters.

## v8.0.0 - 2026-06-01
* feat!: replace BatchActivationSlot with JSON:API relationship model
* Refactor the batch-activation placement API surface to align with the
* JSON:API specification. Slots are no longer embedded directly inside
* `BatchActivationPlacementAttributes`; instead they are exposed via a
* `relationships.slots` to-many relationship on `BatchActivationPlacementData`
* and returned as `batchActivationSlot` entries in the `included` array.
* The `included` field on `PlacementListResponse` and `PlacementResource`
* changes from `[]*ContentStrategyResponse` to `[]*IncludedResource`, a
* new discriminated union that can carry content strategies, batch-activation
* slots, and referenced placements. Slot creation/update fields rename
* `ContentStrategyId` → `PlacementId` throughout.
* Key changes:
* `BatchActivationPlacementAttributes.Slots` field and `GetSlots`/`SetSlots` methods removed; slot detail now lives in `BatchActivationPlacementData.Relationships` (`*BatchActivationPlacementRelationships`)
* `BatchActivationSlot` struct renamed/replaced by `BatchActivationSlotInclusion` (with `BatchActivationSlotAttributes` and `BatchActivationSlotRelationships`); old `SlotId`/`ContentStrategyId` fields removed
* `CreateBatchActivationSlot.ContentStrategyId` renamed to `PlacementId`; `UpdateBatchActivationSlot.ContentStrategyId` renamed to `PlacementId`
* `PlacementListResponse.Included` and `PlacementResource.Included` type changed from `[]*ContentStrategyResponse` to `[]*IncludedResource`
* New exported types added: `BatchActivationPlacementRelationships`, `BatchActivationSlotAttributes`, `BatchActivationSlotInclusion`, `BatchActivationSlotRelationships`, `ContentStrategyInclusion`, `IncludedResource`, `IncludedResourceVisitor`, `PlacementRelationships`, `ResourceIdentifier`, `ToManyRelationship`, `ToOneRelationship`
* 🌿 Generated with Fern

## v7.3.0 - 2026-05-28
### Added
* **`BatchSlotData.Components`** — new optional `*OfferComponents` field exposing slot-level UI components (a `cta` when the slot has no active activation, or a `logoFlare` decoration when it does), accessible via `GetComponents` and `SetComponents`.
* **`BatchSlotData.Assets`** — new optional `[]*Asset` field exposing slot-level visual assets (currently a single `IMG_VIEW` SVG icon), accessible via `GetAssets` and `SetAssets`.

## v7.2.0 - 2026-05-28
### Added
* **`GetEarnedRewardsRequest.FilterPaidInFullOnly`** — new optional `*bool` field (and `SetFilterPaidInFullOnly` setter) that, when `true`, restricts the earned rewards response to transactions paid in full to the issuer (`paidToIssuer` is `PAID_IN_FULL`) and limits `lifetimeRewardsInCents` to those transactions only.

## v7.1.1 - 2026-05-28
* chore: update child organization name validation doc comments
* Update godoc comments across `ChildOrganizationAttributes`,
* `CreateChildAttributes`, `UpdateChildAttributes`, and the `Client.Create`
* method to reflect the revised name validation rule: names must contain at
* least one letter and may include letters and spaces, replacing the previous
* "uppercase, no spaces" constraint.
* Key changes:
* `ChildOrganizationAttributes.Name` doc: "at least one letter; letters and spaces only"
* `CreateChildAttributes.Name` doc: "at least one letter; letters and spaces only"
* `UpdateChildAttributes.Name` doc: "at least one letter; letters and spaces only"
* `Client.Create` godoc: updated inline description to match new name rule
* 🌿 Generated with Fern

## v7.1.0 - 2026-05-27
### Added
* **`EarnedRewardRelationships.Offer`** — new optional `*RelationshipSingle` field representing the offer relationship on an earned reward, accessible via the new `GetOffer` and `SetOffer` methods.

## v7.0.1 - 2026-05-27
* chore: update doc comment links in uploads client
* Update stale relative API documentation links in the `Create` and
* `CreatePart` method godoc comments to use fully-qualified versioned
* paths. No exported signatures, types, or behavior were changed.
* Key changes:
* `Client.Create` godoc: link to Add Upload Part updated from `/api/uploads/create-upload-part` to `/2024-10-01/api/transactions/uploads/create-part`
* `Client.CreatePart` godoc: link to Create Upload updated from `/api/uploads/create-upload` to `/2024-10-01/api/transactions/uploads/create`
* 🌿 Generated with Fern

## v7.0.0 - 2026-05-26
### Breaking Changes
* **`CreateAttributionRequestUnionVisitor`** — a new `VisitPlacementSlotAttribution(*PlacementSlotAttributionRequest) error` method has been added to the interface; all existing implementations must add this method to continue compiling.
### Added
* **`Client.ActivatePlacementSlot`** and **`RawClient.ActivatePlacementSlot`** — new methods that call `POST /v2/issuers/{organizationId}/users/{userId}/placements/{placementId}/slot/{slotId}/activate` to record a slot-level activation event and fan out per-offer ACTIVATE events, returning the event ID and resolved offer IDs.
* **`ActivatePlacementSlotResponse`**, **`ActivatePlacementSlotResponseData`**, and **`ActivatePlacementSlotResponseAttributes`** — new response structs for the slot activation endpoint, exposing `PlacementId`, `SlotId`, `EventCode`, `Medium`, `EventDate`, and `OfferIds`.
* **`PlacementSlotAttributionRequest`**, **`PlacementSlotAttributionAttributes`**, and **`PlacementSlotMedium`** — new types supporting the `placementSlotAttribution` discriminant in `CreateAttributionRequestUnion`, with `PlacementSlotMediumCta` constant and `NewPlacementSlotMediumFromString` constructor.
* **`AttributionState.PlacementId`** and **`AttributionState.SlotId`** — new optional fields (with `GetPlacementId`, `GetSlotId`, `SetPlacementId`, `SetSlotId` methods) to carry placement context on attribution state.

## v6.3.0 - 2026-05-26
### Added
* **`GetBatchesByPlacementRequest`** — new request struct with an optional `SupportedComponents` query parameter for the batch-activation placement endpoint.
* **`BatchSlotData`** — new struct representing a single slot in a batch-activation placement, exposing `SlotId`, `Alias`, `IsActive`, `LastActivatedAt`, `ExpiresAt`, and `Offers` fields with full JSON serialization support.
* **`BatchesResponseObject`** — new response struct wrapping an ordered list of `BatchSlotData` entries returned by the batches endpoint.
* **`Client.PlacementBatches`** and **`RawClient.PlacementBatches`** — new methods that call `GET /v2/issuers/{organizationId}/users/{userId}/placements/{placementId}/batches` to retrieve per-slot offer sets and freshness data for a batch-activation placement.

## v6.2.0 - 2026-05-26
### Added
* **`BatchActivationPlacementAttributes`**, **`BatchActivationPlacementData`**, and **`BatchActivationSlot`** — new structs representing a batch-activation placement and its slots as returned by the API.
* **`CreateBatchActivationAttributes`**, **`CreateBatchActivationPlacementData`**, and **`CreateBatchActivationSlot`** — new structs for supplying batch-activation placement data on create requests.
* **`UpdateBatchActivationAttributes`**, **`UpdateBatchActivationPlacementData`**, and **`UpdateBatchActivationSlot`** — new structs for supplying batch-activation placement data on update requests.
* **`PlacementBatchActivation`** field and **`VisitPlacementBatchActivation`** visitor method on `CreatePlacementDataUnion`, `PlacementFormatUnion`, and `UpdatePlacementDataUnion` — enables handling the new `placementBatchActivation` discriminant in all placement union types.
* **`PlacementTypeFilterPlacementBatchActivation`** — new `PlacementTypeFilter` constant for filtering placements by the batch-activation type.

## v6.1.0 - 2026-05-22
### Added
* **`ProgressBarSegment.Separator`** — new optional field (`*ProgressBarSegmentSeparator`) to control the separator style rendered between segment nodes, with `GetSeparator` and `SetSeparator` methods.
* **`ProgressBarSegment.Labels`** — new optional field (`[]*ProgressBarSegmentLabel`) to supply title/description text for each node in the segment, with `GetLabels` and `SetLabels` methods.
* **`ProgressBarSegment.Selection`** — new optional field (`*ProgressBarSegmentSelection`) to specify which nodes the UI renders as selected, with `GetSelection` and `SetSelection` methods.
* **`ProgressBarSegmentLabel`** — new struct with `Title` and `Description` string fields, full JSON serialization support, and `GetTitle`, `GetDescription`, `SetTitle`, `SetDescription` methods.
* **`ProgressBarSegmentSeparator`** and **`ProgressBarSegmentSelection`** — new string enum types with constants (`LINE`; `CURRENT`, `CURRENT_AND_BELOW`), `NewFrom*` constructors, and `Ptr()` helpers.

## v6.0.0 - 2026-05-21
### Breaking Changes
* **`ContentStrategyAttributes.Filter`**, **`CreateContentStrategyAttributes.Filter`**, and **`UpdateContentStrategyAttributes.Filter`** — the `Filter *ContentStrategyFilter` field has been renamed to `Sort *ContentStrategySort` on all three types. Update direct field access and replace `SetFilter(...)` with `SetSort(...)` and `GetFilter()` with `GetSort()`.
* **`ContentStrategyFilter`** type and its constants (`ContentStrategyFilterNewlyLive`, `ContentStrategyFilterExpiringSoon`, `ContentStrategyFilterHighestCashback`, `ContentStrategyFilterPersonalized`) have been removed. Replace with the equivalent `ContentStrategySort` constants (e.g. `ContentStrategySortNewlyLive`).
* **`NewContentStrategyFilterFromString`** has been removed. Use `NewContentStrategySortFromString` instead.

## v5.0.0 - 2026-05-20
### Breaking Changes
* **`ContentStrategyAttributes.CreatedAt`** and **`ContentStrategyAttributes.LastModified`** — exported fields and their `GetCreatedAt()`, `GetLastModified()`, `SetCreatedAt()`, and `SetLastModified()` methods have been removed. Remove any references to these fields or methods from your code.
* **`MainPagePlacementAttributes.CreatedAt`** and **`MainPagePlacementAttributes.LastModified`** — exported fields and their `GetCreatedAt()`, `GetLastModified()`, `SetCreatedAt()`, and `SetLastModified()` methods have been removed. Remove any references to these fields or methods from your code.
* **`PushNotificationPlacementAttributes.CreatedAt`** and **`PushNotificationPlacementAttributes.LastModified`** — exported fields and their `GetCreatedAt()`, `GetLastModified()`, `SetCreatedAt()`, and `SetLastModified()` methods have been removed. Remove any references to these fields or methods from your code.

## v4.0.0 - 2026-05-20
### Breaking Changes
* **`placements.Client.Get`** now requires a new `*organizations.GetPlacementRequest` parameter and returns `*organizations.PlacementResource` instead of `*organizations.PlacementFormatUnion`. Add `nil` (or a populated `GetPlacementRequest`) as the third argument and update any code that reads the return value directly as a `PlacementFormatUnion` to use `result.GetData()` instead.
* **`placements.RawClient.Get`** now requires a new `*organizations.GetPlacementRequest` parameter and returns `*core.Response[*organizations.PlacementResource]` instead of `*core.Response[*organizations.PlacementFormatUnion]`. Apply the same migration as above.
### Added
* **`GetPlacementRequest`** — new request struct with an optional `Include` field (and `SetInclude` setter) to request sideloaded related resources (e.g. `contentStrategy`) when fetching a single placement.
* **`PlacementResource`** — new response struct wrapping `Data *PlacementFormatUnion` and `Included []*ContentStrategyResponse`, returned by `placements.Client.Get` and `placements.RawClient.Get`.
* **`ListPlacementsRequest.Include`** and **`PlacementListResponse.Included`** — new optional field (with getter/setter) to request and receive sideloaded content strategy resources when listing placements.

## v3.0.0 - 2026-05-20
### Breaking Changes
* **`ContentStrategyAttributes.Filters`**, **`CreateContentStrategyAttributes.Filters`**, and **`UpdateContentStrategyAttributes.Filters`** — the `Filters []ContentStrategyFilter` field has been replaced by `Filter *ContentStrategyFilter` (singular, optional pointer). Update direct field access and replace `SetFilters([]ContentStrategyFilter{...})` with `SetFilter(&value)`.
* **`GetFilters()`** on `ContentStrategyAttributes`, `CreateContentStrategyAttributes`, and `UpdateContentStrategyAttributes` — renamed to **`GetFilter()`** and now returns `*ContentStrategyFilter` instead of `[]ContentStrategyFilter`. Update all call sites accordingly.
* **`SetFilters()`** on `ContentStrategyAttributes`, `CreateContentStrategyAttributes`, and `UpdateContentStrategyAttributes` — renamed to **`SetFilter()`` and now accepts `*ContentStrategyFilter` instead of `[]ContentStrategyFilter`. Update all call sites accordingly.

## v2.2.0 - 2026-05-19
### Added
* **`ListPlacementsRequest.FilterContentStrategyId`** — new optional filter field (with `SetFilterContentStrategyId` setter) to narrow placement listings by the ID of a linked content strategy.
* **`CreateMainPageAttributes.ContentStrategyId`** and **`CreatePushNotificationAttributes.ContentStrategyId`** — new optional field (with getter/setter) to link a content strategy to a placement at creation time.
* **`UpdateMainPageAttributes.ContentStrategyId`** and **`UpdatePushNotificationAttributes.ContentStrategyId`** — new optional field (with getter/setter) to link or unlink a content strategy when updating a placement.
* **`MainPagePlacementAttributes.ContentStrategyId`** and **`PushNotificationPlacementAttributes.ContentStrategyId`** — new optional response field (with getter/setter) exposing the ID of the content strategy linked to a placement.

## v2.1.0 - 2026-05-19
### Added
* **`contentstrategies.Client`** — new sub-client on the organizations `Client` with `Create`, `List`, `Get`, `Update`, and `Delete` methods for managing content strategies scoped to an organization.
* **`contentstrategies.RawClient`** — companion raw client returning full HTTP response metadata alongside typed response bodies for all content strategy operations.
* **`ContentStrategyAttributes`**, **`ContentStrategyResponse`**, and **`ContentStrategyListResponse`** — new structs representing content strategy resources with full getter/setter and JSON serialization support.
* **`CreateContentStrategyRequestBody`**, **`UpdateContentStrategyRequestBody`**, and **`ListContentStrategiesRequest`** — new request types for creating, updating, and paginated listing of content strategies.
* **`ContentStrategyFilter`** — new string enum type with constants for `NewlyLive`, `ExpiringSoon`, `HighestCashback`, and `Personalized` filters, plus a `NewContentStrategyFilterFromString` constructor.

## v2.0.1 - 2026-05-14
* chore: update LocationAttributes.PartnerIds comment and omitempty tag
* Refine the `PartnerIds` field on `LocationAttributes` by updating its
* godoc comment from "Only included on LOCAL locations" to "Only applicable
* for LOCAL locations" and removing the `omitempty` JSON/URL tag so the
* field is always serialized when present.
* Key changes:
* Update `PartnerIds` godoc comment for clarity ("Only applicable for LOCAL locations")
* Remove `omitempty` from `json` and `url` struct tags on `PartnerIds`
* 🌿 Generated with Fern

## v2.0.0 - 2026-05-14
### Breaking Changes
* **`EarnedRewardAttributes`** — struct and all associated getter/setter methods have been removed; migrate by replacing any usage with `RewardNotificationAttributes`.
* **`EarnedRewardApprovedData.Attributes`** — field type changed from `*EarnedRewardAttributes` to `*RewardNotificationAttributes`; update all call sites to use `RewardNotificationAttributes`.
### Added
* **`RewardNotificationAttributes.TransactionId`**, **`TransactionAmountInCents`**, and **`TransactionTimestamp`** — new fields (with corresponding getters and setters) exposing originating transaction details on reward notification attributes.
* **`EarnedRewardSettledAttributes.TransactionId`** and **`TransactionAmountInCents`** — new fields providing transaction ID and amount in cents for settled reward events.
* **`ValidTransactionAttributes.TransactionId`**, **`TransactionAmountInCents`**, and **`TransactionTimestamp`** — new fields providing originating transaction details for valid transaction notification attributes.

## v1.1.0 - 2026-05-12
### Added
* **`LocationPartnerId`** — new struct representing a third-party partner ID (e.g. a Google place ID) associated with a location.
* **`LocationPartnerIdType`** — new string enum type with the `LocationPartnerIdTypeGoogle` constant and a `NewLocationPartnerIdTypeFromString` constructor.
* **`LocationAttributes.PartnerIds`** — new optional field (`[]*LocationPartnerId`) exposing partner IDs for LOCAL locations, accessible via `GetPartnerIds` and `SetPartnerIds`.

## v1.0.1 - 2026-05-07
* fix: correct required scope in CreateBulkTransactionsUploadUrl doc
* Update the documented required OAuth scope for the
* `CreateBulkTransactionsUploadUrl` method from `transaction:write` to
* `files:write`, reflecting the actual permission required by the API.
* Key changes:
* Fix godoc comment on `CreateBulkTransactionsUploadUrl` to reference
* the correct required scope (`files:write` instead of `transaction:write`)
* 🌿 Generated with Fern

## 1.0.0 - 2026-05-05
* Initial release of the Go SDK.
* Provides core client functionality, authentication support, and documented usage examples to help you get started quickly.
