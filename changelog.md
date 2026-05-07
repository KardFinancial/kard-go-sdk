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
