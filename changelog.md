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
