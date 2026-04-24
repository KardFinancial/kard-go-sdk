# Reference
## Auth
<details><summary><code>client.Auth.GetToken(request) -> *kardgosdk.TokenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.GetTokenRequest{
        ClientID: "client_id",
        ClientSecret: "client_secret",
    }
client.Auth.GetToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**xKardTargetIssuer:** `*string` — (Beta) Target issuer ID for partners managing multiple issuers on the Kard platform. When set, the auth token will be scoped to this specific issuer. Required if you manage more than one issuer; omit if you operate a single issuer integration.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Files
<details><summary><code>client.Files.GetMetadata(OrganizationID) -> *kardgosdk.GetFilesMetadataResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves metadata for files associated with a specific issuer/organization.
This endpoint supports pagination and sorting options to efficiently navigate 
through potentially large sets of file metadata.
<b>Required scopes:</b> `files.read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.GetFilesMetadataRequest{
        PageSize: kardgosdk.Int(
            5,
        ),
        FilterDateFrom: kardgosdk.String(
            "2025-02-20T21:56:23Z",
        ),
        FilterDateTo: kardgosdk.String(
            "2025-03-20T21:56:23Z",
        ),
        FilterFileType: kardgosdk.FileTypeEarnedRewardApprovedDailyReconciliationFile.Ptr(),
        Sort: []*kardgosdk.FilesMetadataSortOptions{
            kardgosdk.FilesMetadataSortOptionsSentDateDesc.Ptr(),
        },
    }
client.Files.GetMetadata(
        context.TODO(),
        "organization-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**filterDateFrom:** `*string` — Start date for filtering files (format ISO8601). If not provided, defaults to current date minus 1 month.
    
</dd>
</dl>

<dl>
<dd>

**filterDateTo:** `*string` — End date for filtering files (format ISO8601). If not provided, defaults to current date.
    
</dd>
</dl>

<dl>
<dd>

**filterFileType:** `*kardgosdk.FileType` — The document file type.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of items per page. Defaults to 10 if not specified and maximum value allowed 100 items per page.
    
</dd>
</dl>

<dl>
<dd>

**pageAfter:** `*string` — Cursor for forward pagination (next page).
    
</dd>
</dl>

<dl>
<dd>

**pageBefore:** `*string` — Cursor for backward pagination (previous page).
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*kardgosdk.FilesMetadataSortOptions` — If provided, response will be sorted by the specified fields. Defaults to descending sentDate, equivalent to "-sentDate"
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Subscriptions
<details><summary><code>client.Notifications.Subscriptions.Get(OrganizationID) -> *notifications.SubscriptionsResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to fetch the subscriptions of the provided issuer.<br/>
<b>Required scopes:</b> `notifications:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.GetSubscriptionsRequest{}
client.Notifications.Subscriptions.Get(
        context.TODO(),
        "organization-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**filterEventName:** `*kardgosdk.NotificationType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.Subscriptions.Create(OrganizationID, request) -> *notifications.CreateSubscriptionsResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to subscribe to notification events.<br/>
<b>Required scopes:</b> `notifications:write`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.SubscriptionRequestBody{
        Data: []*notifications.SubscriptionRequestUnion{
            &notifications.SubscriptionRequestUnion{
                Subscription: &notifications.SubscriptionRequest{
                    Attributes: &notifications.SubscriptionRequestAttributes{
                        EventName: kardgosdk.NotificationTypeEarnedRewardApproved,
                        WebhookURL: "https://webhookUrl.com/post",
                        Enabled: true,
                    },
                },
            },
        },
    }
client.Notifications.Subscriptions.Create(
        context.TODO(),
        "organization-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*notifications.SubscriptionRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.Subscriptions.Update(OrganizationID, SubscriptionID, request) -> *notifications.UpdateSubscriptionsResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to update existing notification subscriptions.<br/>
<b>Required scopes:</b> `notifications:write`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &notifications.UpdateSubscriptionRequestBody{
        Data: &notifications.UpdateSubscriptionRequestUnion{
            Subscription: &notifications.UpdateSubscriptionRequest{
                Attributes: &notifications.UpdateSubscriptionRequestAttributes{
                    EventName: kardgosdk.NotificationTypeEarnedRewardApproved.Ptr(),
                    WebhookURL: kardgosdk.String(
                        "https://webhookUrl.com/post",
                    ),
                    Enabled: kardgosdk.Bool(
                        true,
                    ),
                },
            },
        },
    }
client.Notifications.Subscriptions.Update(
        context.TODO(),
        "organization-123",
        "subscription-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**subscriptionID:** `kardgosdk.SubscriptionID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*notifications.UpdateSubscriptionRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations
<details><summary><code>client.Organizations.Get() -> *kardgosdk.ExternalOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve organization details for the authenticated issuer
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Children
<details><summary><code>client.Organizations.Children.List(OrganizationID) -> *organizations.ChildOrganizationListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List child organizations belonging to the authenticated issuer
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &organizations.ListChildrenRequest{}
client.Organizations.Children.List(
        context.TODO(),
        "organizationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**pageAfter:** `*string` — Cursor value for the next page of results
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Maximum number of records to return [1 - 200] (default = 200)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Children.Create(OrganizationID, request) -> *organizations.ChildOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a child organization by cloning the parent and overriding specified fields. An 8-digit numeric ID is generated automatically. The name is required, must be uppercase, and must not contain spaces.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &organizations.CreateChildRequestBody{
        Data: &organizations.CreateChildRequestData{
            Attributes: &organizations.CreateChildAttributes{
                Name: "name",
            },
        },
    }
client.Organizations.Children.Create(
        context.TODO(),
        "organizationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**request:** `*organizations.CreateChildRequestBody` — Child organization data for creation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Children.Get(OrganizationID, ChildID) -> *organizations.ChildOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific child organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Children.Get(
        context.TODO(),
        "organizationId",
        "childId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**childID:** `string` — Unique identifier of the child organization
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Children.Update(OrganizationID, ChildID, request) -> *organizations.ChildOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a child organization. Only the name can be changed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &organizations.UpdateChildRequestBody{
        Data: &organizations.UpdateChildRequestData{
            Attributes: &organizations.UpdateChildAttributes{},
        },
    }
client.Organizations.Children.Update(
        context.TODO(),
        "organizationId",
        "childId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**childID:** `string` — Unique identifier of the child organization
    
</dd>
</dl>

<dl>
<dd>

**request:** `*organizations.UpdateChildRequestBody` — Child organization data for update
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Children.Delete(OrganizationID, ChildID) -> *kardgosdk.DeleteResourceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a child organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Children.Delete(
        context.TODO(),
        "organizationId",
        "childId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**childID:** `string` — Unique identifier of the child organization
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Placements
<details><summary><code>client.Organizations.Placements.Create(OrganizationID, request) -> *organizations.PlacementFormatUnion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a placement for the organization. Use type "placementMainPage" for main-page placements (requires name and availableSlots) or "placementPushNotification" for push-notification placements (requires name and cadence; availableSlots is automatically set to 1).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &organizations.CreatePlacementRequestBody{
        Data: &organizations.CreatePlacementDataUnion{
            PlacementMainPage: &organizations.CreateMainPagePlacementData{
                Attributes: &organizations.CreateMainPageAttributes{
                    Name: "Homepage Banner",
                    AvailableSlots: 5,
                },
            },
        },
    }
client.Organizations.Placements.Create(
        context.TODO(),
        "org-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**request:** `*organizations.CreatePlacementRequestBody` — Placement data for creation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Placements.List(OrganizationID) -> *organizations.PlacementListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List placements belonging to the authenticated organization
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &organizations.ListPlacementsRequest{}
client.Organizations.Placements.List(
        context.TODO(),
        "organizationId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**filterType:** `*organizations.PlacementTypeFilter` — Filter by placement type (placementMainPage or placementPushNotification)
    
</dd>
</dl>

<dl>
<dd>

**filterName:** `*string` — Filter by exact placement name (unique within an organization per type)
    
</dd>
</dl>

<dl>
<dd>

**pageAfter:** `*string` — Cursor value for the next page of results
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Maximum number of records to return [1 - 200] (default = 200)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Placements.Get(OrganizationID, PlacementID) -> *organizations.PlacementFormatUnion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific placement
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Placements.Get(
        context.TODO(),
        "organizationId",
        "placementId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**placementID:** `string` — Unique identifier of the placement (UUID v7)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Placements.Update(OrganizationID, PlacementID, request) -> *organizations.PlacementFormatUnion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace a placement. All fields must be provided. Use type "placementMainPage" or "placementPushNotification" to set the placement kind. If the type is "placementPushNotification", availableSlots is automatically set to 1.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &organizations.UpdatePlacementRequestBody{
        Data: &organizations.UpdatePlacementDataUnion{
            PlacementMainPage: &organizations.UpdateMainPagePlacementData{
                Attributes: &organizations.UpdateMainPageAttributes{
                    Name: "name",
                    AvailableSlots: 1,
                },
            },
        },
    }
client.Organizations.Placements.Update(
        context.TODO(),
        "organizationId",
        "placementId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**placementID:** `string` — Unique identifier of the placement (UUID v7)
    
</dd>
</dl>

<dl>
<dd>

**request:** `*organizations.UpdatePlacementRequestBody` — Placement data for update
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Placements.Delete(OrganizationID, PlacementID) -> *kardgosdk.DeleteResourceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a placement
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Placements.Delete(
        context.TODO(),
        "organizationId",
        "placementId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**placementID:** `string` — Unique identifier of the placement (UUID v7)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ping
<details><summary><code>client.Ping.Ping() -> *kardgosdk.PingResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to verify network connectivity and service availability.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ping.Ping(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Transactions
<details><summary><code>client.Transactions.Create(OrganizationID, request) -> *kardgosdk.TransactionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to send all transactions made by all your enrolled users in your rewards program. The request body will depend on the transaction type.<br/>
Please use the correct type when calling the endpoint:
- `transaction`: These incoming transactions will be processed and matched by the Kard system. Learn more about the [Transaction CLO Matching](https://github.com/kard-financial/kard-postman#c-transaction-clo-matching) flow here.
- `matchedTransaction`: For pre-matched transactions that need validation on match by the Kard system.
- `coreTransaction`: For transactions from core banking systems with limited card-level data.<br/>

<b>Required scopes:</b> `transaction:write`<br/>
<b>Note:</b> `Maximum of 500 transactions can be created per request`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.TransactionsRequestBody{
        Data: []*kardgosdk.Transactions{
            &kardgosdk.Transactions{
                Transaction: &kardgosdk.TransactionsRequest{
                    ID: "309rjfoincor3icno3rind093cdow3jciwjdwcm",
                    Attributes: &kardgosdk.TransactionsAttributes{
                        UserID: "6FHt5b6Fnp0qdomMEy5AN6PXcSJIeX69",
                        Status: kardgosdk.TransactionStatusApproved,
                        Amount: 1000,
                        Subtotal: kardgosdk.Int(
                            800,
                        ),
                        Currency: "USD",
                        Direction: kardgosdk.DirectionTypeDebit,
                        PaymentType: kardgosdk.TransactionPaymentTypeCard,
                        Description: "ADVANCEAUTO",
                        Description2: kardgosdk.String(
                            "ADVANCEAUTO",
                        ),
                        Mcc: kardgosdk.String(
                            "1234",
                        ),
                        CardBin: "123456",
                        CardLastFour: "4321",
                        AuthorizationDate: kardgosdk.Time(
                            kardgosdk.MustParseDateTime(
                                "2021-07-02T17:47:06Z",
                            ),
                        ),
                        Merchant: &kardgosdk.Merchant{
                            ID: kardgosdk.String(
                                "12345678901234567",
                            ),
                            Name: "ADVANCEAUTO",
                            AddrStreet: kardgosdk.String(
                                "125 Main St",
                            ),
                            AddrCity: kardgosdk.String(
                                "Philadelphia",
                            ),
                            AddrState: kardgosdk.StatesPa.Ptr(),
                            AddrZipcode: kardgosdk.String(
                                "19147",
                            ),
                            AddrCountry: kardgosdk.String(
                                "United States",
                            ),
                            Latitude: kardgosdk.String(
                                "37.9419429",
                            ),
                            Longitude: kardgosdk.String(
                                "-73.1446869",
                            ),
                            StoreID: kardgosdk.String(
                                "12345",
                            ),
                        },
                        AuthorizationCode: kardgosdk.String(
                            "123456",
                        ),
                        RetrievalReferenceNumber: kardgosdk.String(
                            "100804333919",
                        ),
                        AcquirerReferenceNumber: kardgosdk.String(
                            "1234567890123456789012345678",
                        ),
                        SystemTraceAuditNumber: kardgosdk.String(
                            "333828",
                        ),
                        TransactionID: "2467de37-cbdc-416d-a359-75de87bfffb0",
                        CardProductID: kardgosdk.String(
                            "1234567890123456789012345678",
                        ),
                        ProcessorMids: &kardgosdk.ProcessorMid{
                            Visa: &kardgosdk.VisaMid{
                                Mids: &kardgosdk.VisaMidDetails{
                                    Vmid: "12345678901",
                                    Vsid: "12345678",
                                },
                            },
                        },
                    },
                },
            },
        },
    }
client.Transactions.Create(
        context.TODO(),
        "organization-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kardgosdk.TransactionsRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.CreateFraudMarkers(OrganizationID, request) -> *kardgosdk.FraudulentTransactionObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to flag a submitted transaction as fraudulent. This will prevent it from being rewarded.<br/>

<b>Required scopes:</b>&nbsp;&nbsp;`transaction:write`<br/>
<b>Note:</b> `Maximum of 500 fraudulent transactions can be created per request`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.FraudulentTransactionRequestBody{
        Data: []*kardgosdk.FraudulentTransactionData{
            &kardgosdk.FraudulentTransactionData{
                ID: "myTxnId12345",
                Type: "fraudulentTransaction",
                Attributes: &kardgosdk.FraudulentTransactionAttributes{
                    UserID: "userId123",
                },
            },
        },
    }
client.Transactions.CreateFraudMarkers(
        context.TODO(),
        "organization-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kardgosdk.FraudulentTransactionRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.CreateAudits(OrganizationID, UserID, request) -> *kardgosdk.CreateAuditResponseBody</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to request that a particular transaction be audited further by the Kard system, in the event of a missing cashback claim, incorrect cashback amount claim or other mis-match claims.<br/>
<b>Required scopes:</b> `audit:write`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.CreateAuditRequestBody{
        Data: []*kardgosdk.CreateAuditRequestDataUnion{
            &kardgosdk.CreateAuditRequestDataUnion{
                Audit: &kardgosdk.AuditRequestData{
                    Attributes: &kardgosdk.AuditAttributes{
                        AuditCode: 8001,
                        MerchantName: "Caribbean Goodness",
                        AuditDescription: "duplicate transaction",
                        TransactionID: "issuerTransaction123",
                    },
                },
            },
        },
    }
client.Transactions.CreateAudits(
        context.TODO(),
        "organization-123",
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kardgosdk.CreateAuditRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.CreateBulkTransactionsUploadURL(OrganizationID, request) -> *kardgosdk.CreateFileUploadURLResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Generates up to 10 presigned PUT URLs for uploading JSONL transaction files (up to 5GB each) directly
to storage. Each URL is valid for 15 minutes. Use the returned URL to upload the file via an HTTP PUT request with the
binary file content as the body. If a URL expires before the upload completes, you must request a new one.
Files can be uploaded as plain JSONL or as a gzip-compressed file.
Only `coreTransaction` type is supported for bulk file uploads.
<b>Required scopes:</b> `transaction:write`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.CreateFileUploadRequestBody{
        Data: []*kardgosdk.CreateFileUploadData{
            &kardgosdk.CreateFileUploadData{
                Type: kardgosdk.FileUploadTypeIncomingTransactionsFile,
                Attributes: &kardgosdk.CreateFileUploadAttributes{
                    Filename: "transaction_12345.jsonl",
                },
            },
            &kardgosdk.CreateFileUploadData{
                Type: kardgosdk.FileUploadTypeIncomingTransactionsFile,
                Attributes: &kardgosdk.CreateFileUploadAttributes{
                    Filename: "transaction_67890.jsonl",
                },
            },
        },
    }
client.Transactions.CreateBulkTransactionsUploadURL(
        context.TODO(),
        "organization-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kardgosdk.CreateFileUploadRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.GetEarnedRewards(OrganizationID, UserID) -> *kardgosdk.GetEarnedRewardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve rewarded transaction history for a specific user. By default this returns only SETTLED transactions within the last 12 months.
<br/>
<b>Required scopes:</b> `transaction:read`
<br/>
<b>Query Limit:</b> Maximum of 12 months of transaction data can be queried.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.GetEarnedRewardsRequest{
        PageSize: kardgosdk.Int(
            10,
        ),
        FilterStatus: kardgosdk.RewardedTransactionStatusApproved.Ptr(),
        Include: kardgosdk.String(
            "merchant,offer",
        ),
    }
client.Transactions.GetEarnedRewards(
        context.TODO(),
        "org-123",
        "user-456",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**pageAfter:** `*string` — Cursor for next page (base64-encoded timestamp + transaction ID)
    
</dd>
</dl>

<dl>
<dd>

**pageBefore:** `*string` — Cursor for previous page (base64-encoded timestamp + transaction ID)
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results per page
    
</dd>
</dl>

<dl>
<dd>

**filterStatus:** `*kardgosdk.RewardedTransactionStatus` — Filter by transaction status. Supported values are `APPROVED` and `SETTLED`. Defaults to `SETTLED` when omitted. When `APPROVED` is specified, only approved transactions that do not yet have a corresponding settled transaction are returned.
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — Comma-separated list of related resources to include in the response. Supported values are `merchant` and `offer`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users
<details><summary><code>client.Users.Create(OrganizationID, request) -> *kardgosdk.CreateUsersObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to enroll a specified user into your rewards program.<br/>

<b>Required scopes:</b>&nbsp;&nbsp;`user:write`<br/>
<b>Note:</b> `Maximum of 100 users can be created per request`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.CreateUsersObject{
        Data: []*kardgosdk.UserRequestDataUnion{
            &kardgosdk.UserRequestDataUnion{
                User: &kardgosdk.UserRequestData{
                    ID: "1234567890",
                    Attributes: &kardgosdk.UserRequestAttributes{
                        ZipCode: kardgosdk.String(
                            "11238",
                        ),
                        EnrolledRewards: []kardgosdk.EnrolledRewardsType{
                            kardgosdk.EnrolledRewardsTypeCardlinked,
                        },
                        Email: kardgosdk.String(
                            "user@example.com",
                        ),
                        HashedEmail: kardgosdk.String(
                            "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3e2d8a5b76e45a1d4c4e2e3a1",
                        ),
                        PhoneNumber: kardgosdk.String(
                            "+14155552671",
                        ),
                        BirthYear: kardgosdk.String(
                            "1990",
                        ),
                        HistoricalTransactionsSent: kardgosdk.Bool(
                            true,
                        ),
                    },
                },
            },
        },
    }
client.Users.Create(
        context.TODO(),
        "organization-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kardgosdk.CreateUsersObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Update(OrganizationID, UserID, request) -> *kardgosdk.UserResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to update the details on a specified user.<br/>

<b>Required scopes:</b> `user:update`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kardgosdk.UpdateUserObject{
        Data: &kardgosdk.UpdateUserRequestDataUnion{
            User: &kardgosdk.UpdateUserRequestData{
                ID: "1234567890",
                Attributes: &kardgosdk.UpdateUserRequestAttributes{
                    ZipCode: kardgosdk.String(
                        "11238",
                    ),
                    EnrolledRewards: []kardgosdk.EnrolledRewardsType{
                        kardgosdk.EnrolledRewardsTypeCardlinked,
                    },
                    Email: kardgosdk.String(
                        "user@example.com",
                    ),
                    HashedEmail: kardgosdk.String(
                        "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3e2d8a5b76e45a1d4c4e2e3a1",
                    ),
                    PhoneNumber: kardgosdk.String(
                        "+14155552671",
                    ),
                    BirthYear: kardgosdk.String(
                        "1990",
                    ),
                },
            },
        },
    }
client.Users.Update(
        context.TODO(),
        "organization-123",
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kardgosdk.UpdateUserObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Delete(OrganizationID, UserID) -> *kardgosdk.DeleteUserResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to delete a specified enrolled user from the rewards program and Kard's system. Users can be re-enrolled into rewards by calling the [Create User](/2024-10-01/api/users/create) endpoint using the same `id` from before.<br/>

<b>Required scopes:</b> `user:delete`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Delete(
        context.TODO(),
        "organization-123",
        "user-123",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Get(OrganizationID, UserID) -> *kardgosdk.UserResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to fetch the details on a specified user.<br/>
<br/>
<b>Required scopes:</b>&nbsp;&nbsp;`user:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Get(
        context.TODO(),
        "organization-123",
        "user-123",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Attributions
<details><summary><code>client.Users.Attributions.Create(OrganizationID, UserID, request) -> *users.CreateAttributionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to send attribution events made by a single enrolled user for processing. A maximum of 100 events can be included in a single request.

<b>Required scopes:</b> `attributions:write`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.CreateAttributionRequestObject{
        Data: []*users.CreateAttributionRequestUnion{
            &users.CreateAttributionRequestUnion{
                OfferAttribution: &users.OfferAttributionRequest{
                    Attributes: &users.OfferAttributionAttributes{
                        EntityID: "60e4ba1da31c5a22a144c075",
                        EventCode: users.EventCodeView,
                        Medium: users.OfferMediumSearch,
                        EventDate: kardgosdk.MustParseDateTime(
                            "2025-01-01T00:00:00Z",
                        ),
                    },
                },
            },
            &users.CreateAttributionRequestUnion{
                OfferAttribution: &users.OfferAttributionRequest{
                    Attributes: &users.OfferAttributionAttributes{
                        EntityID: "60e4ba1da31c5a22a144c077",
                        EventCode: users.EventCodeImpression,
                        Medium: users.OfferMediumEmail,
                        EventDate: kardgosdk.MustParseDateTime(
                            "2025-01-01T00:00:00Z",
                        ),
                    },
                },
            },
            &users.CreateAttributionRequestUnion{
                NotificationAttribution: &users.NotificationAttributionRequest{
                    Attributes: &users.NotificationAttributionAttributes{
                        EntityID: "60e4ba1da31c5a22a144c076",
                        EventCode: users.EventCodeImpression,
                        Medium: users.NotificationMediumPush,
                        EventDate: kardgosdk.MustParseDateTime(
                            "2025-01-01T00:00:00Z",
                        ),
                    },
                },
            },
        },
    }
client.Users.Attributions.Create(
        context.TODO(),
        "organization-123",
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*users.CreateAttributionRequestObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Attributions.Activate(OrganizationID, UserID, OfferID) -> *users.ActivateOfferResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Record when a user activates an offer. Creates an attribution event with eventCode=ACTIVATE and medium=CTA.
Optionally include the offer data by passing `include=offer`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.ActivateOfferRequest{}
client.Users.Attributions.Activate(
        context.TODO(),
        "organization-123",
        "user-123",
        "offer-456",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**offerID:** `string` — The unique identifier of the offer being activated
    
</dd>
</dl>

<dl>
<dd>

**supportedComponents:** `*users.ComponentType` — UI component types to include in the offer response (when include=offer).
    
</dd>
</dl>

<dl>
<dd>

**include:** `*users.ActivateOfferIncludeOption` — Related resources to include in the response. Allowed value is `offer`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Attributions.Boost(OrganizationID, UserID, OfferID) -> *users.BoostOfferResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Record when a user boosts an offer. Creates an attribution event with eventCode=BOOST and medium=CTA.
Optionally include the offer data by passing `include=offer`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.BoostOfferRequest{}
client.Users.Attributions.Boost(
        context.TODO(),
        "organization-123",
        "user-123",
        "offer-456",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**offerID:** `string` — The unique identifier of the offer being boosted
    
</dd>
</dl>

<dl>
<dd>

**supportedComponents:** `*users.ComponentType` — UI component types to include in the offer response (when include=offer).
    
</dd>
</dl>

<dl>
<dd>

**include:** `*users.BoostOfferIncludeOption` — Related resources to include in the response. Allowed value is `offer`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## WebView
<details><summary><code>client.Users.Auth.GetWebViewToken(OrganizationID, UserID) -> *users.WebViewTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves an OAuth token for webview authentication.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Auth.GetWebViewToken(
        context.TODO(),
        "organization-123",
        "user-123",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Rewards
<details><summary><code>client.Users.Rewards.Offers(OrganizationID, UserID) -> *users.OffersResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve national brand offers that a specified user is eligible for. Call this endpoint to build out your
[targeted offers UX experience](/2024-10-01/api/getting-started#b-discover-a-lapsed-customer-clo). Local offers details
can be found by calling the [Get Eligible Locations](/2024-10-01/api/rewards/locations).<br/>
<b>Required scopes:</b> `rewards:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.GetOffersByUserRequest{
        PageSize: kardgosdk.Int(
            1,
        ),
        FilterIsTargeted: kardgosdk.Bool(
            true,
        ),
        Sort: []*users.OfferSortOptions{
            users.OfferSortOptionsStartDateDesc.Ptr(),
        },
    }
client.Users.Rewards.Offers(
        context.TODO(),
        "organization-123",
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pageBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**filterSearch:** `*string` — Case-insensitive search string to filter offers by merchant name
    
</dd>
</dl>

<dl>
<dd>

**filterPurchaseChannel:** `[]*kardgosdk.PurchaseChannel` 
    
</dd>
</dl>

<dl>
<dd>

**filterCategory:** `*kardgosdk.CategoryOption` 
    
</dd>
</dl>

<dl>
<dd>

**filterIsTargeted:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*users.OfferSortOptions` — If provided, response will be sorted by the specified fields
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — CSV list of included resources in the response (e.g "categories"). Allowed value is `categories`.
    
</dd>
</dl>

<dl>
<dd>

**supportedComponents:** `*users.ComponentType` — UI component types to include in the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Rewards.PlacementOffers(OrganizationID, UserID, PlacementID) -> *users.OffersResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve offers for a placement slot. Returns offers sorted by highest cash back,
limited by the placement's available slots.<br/>
<b>Required scopes:</b> `rewards:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.GetOffersByPlacementRequest{}
client.Users.Rewards.PlacementOffers(
        context.TODO(),
        "organizationId",
        "userId",
        "placementId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**placementID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**filterSearch:** `*string` — Case-insensitive search string to filter offers by merchant name
    
</dd>
</dl>

<dl>
<dd>

**filterPurchaseChannel:** `[]*kardgosdk.PurchaseChannel` 
    
</dd>
</dl>

<dl>
<dd>

**filterCategory:** `*kardgosdk.CategoryOption` 
    
</dd>
</dl>

<dl>
<dd>

**filterIsTargeted:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — CSV list of included resources in the response (e.g "categories"). Allowed value is `categories`.
    
</dd>
</dl>

<dl>
<dd>

**supportedComponents:** `*users.ComponentType` — UI component types to include in the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Rewards.Locations(OrganizationID, UserID) -> *users.LocationsResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve national and local geographic locations that a specified user has eligible in-store offers at. Use this endpoint to build
out your [map-specific UX experiences](/2024-10-01/api/getting-started#c-discover-clos-near-you-map-view). Please note
that Longitude and Latitude fields are prioritized over State, City and Zipcode and are the recommended search
pattern.<br/>
<br/>
<b>Required scopes:</b> `rewards:read`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.GetLocationsByUserRequest{
        PageSize: kardgosdk.Int(
            1,
        ),
        FilterLatitude: kardgosdk.Float64(
            39.9419429,
        ),
        FilterLongitude: kardgosdk.Float64(
            -75.1446869,
        ),
        FilterRadius: kardgosdk.Int(
            10,
        ),
        Include: []*string{
            kardgosdk.String(
                "offers,categories",
            ),
        },
    }
client.Users.Rewards.Locations(
        context.TODO(),
        "organization-123",
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `kardgosdk.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pageBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**filterName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**filterCity:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**filterZipCode:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**filterState:** `*kardgosdk.State` 
    
</dd>
</dl>

<dl>
<dd>

**filterCategory:** `*kardgosdk.CategoryOption` 
    
</dd>
</dl>

<dl>
<dd>

**filterLongitude:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**filterLatitude:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**filterRadius:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*users.LocationSortOptions` — If provided, response will be sorted by the specified fields
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — CSV list of included resources in the response (e.g "offers,categories"). Allowed values are `offers` and `categories`.
    
</dd>
</dl>

<dl>
<dd>

**supportedComponents:** `*users.ComponentType` — UI component types to include in included offers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Uploads
<details><summary><code>client.Users.Uploads.Create(OrganizationID, UserID, request) -> *users.CreateUploadResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to create an upload session and retrieve an upload ID. Using the upload ID in the [Add Upload 
Part](/api/uploads/create-upload-part) endpoint, historical transactions can be sent in batches for further processing.
<b>Required scopes:</b> `transaction:write`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.CreateUploadRequestObject{
        Data: &users.CreateUploadRequestDataUnion{
            HistoricalTransactionStart: &users.StartHistoricalUploadNoData{
                Attributes: &kardgosdk.EmptyObject{},
            },
        },
    }
client.Users.Uploads.Create(
        context.TODO(),
        "organization-123",
        "user-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**request:** `*users.CreateUploadRequestObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Uploads.CreatePart(OrganizationID, UserID, UploadID, request) -> *users.CreateUploadPartResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint using the upload ID provided in the [Create Upload](/api/uploads/create-upload) endpoint to add parts to your upload. Currently, this endpoint supports adding historical transactions.
<b>Required scopes:</b> `transaction:write`
<b>Note:</b> `Maximum of 500 transactions can be uploaded per request`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.CreateUploadPartRequestObject{
        Data: []*users.CreateUploadPartDataUnion{
            &users.CreateUploadPartDataUnion{
                HistoricalTransaction: &kardgosdk.TransactionsRequest{
                    ID: "309rjfoincor3icno3rind093cdow3jciwjdwcm",
                    Attributes: &kardgosdk.TransactionsAttributes{
                        UserID: "6FHt5b6Fnp0qdomMEy5AN6PXcSJIeX69",
                        Status: kardgosdk.TransactionStatusApproved,
                        Amount: 1000,
                        Subtotal: kardgosdk.Int(
                            800,
                        ),
                        Currency: "USD",
                        Direction: kardgosdk.DirectionTypeDebit,
                        PaymentType: kardgosdk.TransactionPaymentTypeCard,
                        Description: "ADVANCEAUTO",
                        Description2: kardgosdk.String(
                            "ADVANCEAUTO",
                        ),
                        Mcc: kardgosdk.String(
                            "1234",
                        ),
                        CardBin: "123456",
                        CardLastFour: "4321",
                        AuthorizationDate: kardgosdk.Time(
                            kardgosdk.MustParseDateTime(
                                "2021-07-02T17:47:06Z",
                            ),
                        ),
                        Merchant: &kardgosdk.Merchant{
                            ID: kardgosdk.String(
                                "12345678901234567",
                            ),
                            Name: "ADVANCEAUTO",
                            AddrStreet: kardgosdk.String(
                                "125 Main St",
                            ),
                            AddrCity: kardgosdk.String(
                                "Philadelphia",
                            ),
                            AddrState: kardgosdk.StatesPa.Ptr(),
                            AddrZipcode: kardgosdk.String(
                                "19147",
                            ),
                            AddrCountry: kardgosdk.String(
                                "United States",
                            ),
                            Latitude: kardgosdk.String(
                                "37.9419429",
                            ),
                            Longitude: kardgosdk.String(
                                "-73.1446869",
                            ),
                            StoreID: kardgosdk.String(
                                "12345",
                            ),
                        },
                        AuthorizationCode: kardgosdk.String(
                            "123456",
                        ),
                        RetrievalReferenceNumber: kardgosdk.String(
                            "100804333919",
                        ),
                        AcquirerReferenceNumber: kardgosdk.String(
                            "1234567890123456789012345678",
                        ),
                        SystemTraceAuditNumber: kardgosdk.String(
                            "333828",
                        ),
                        TransactionID: "2467de37-cbdc-416d-a359-75de87bfffb0",
                    },
                },
            },
        },
    }
client.Users.Uploads.CreatePart(
        context.TODO(),
        "organization-123",
        "user-123",
        "upload-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**uploadID:** `string` — The upload ID identifying the upload session to add parts
    
</dd>
</dl>

<dl>
<dd>

**request:** `*users.CreateUploadPartRequestObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Uploads.Update(OrganizationID, UserID, UploadID, request) -> *users.UpdateUploadResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Call this endpoint to update your upload session. Currently, you can signal completing a historical transactions upload.
<b>Required scopes:</b> `transaction:write`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &users.UpdateUploadRequestObject{
        Data: &users.UpdateUploadRequestDataUnion{
            HistoricalTransactionComplete: &users.HistoricalTransactionCompleteNoData{
                ID: "7e382223-b9a5-4825-91fb-436c8957a2e7",
                Attributes: &kardgosdk.EmptyObject{},
            },
        },
    }
client.Users.Uploads.Update(
        context.TODO(),
        "organization-123",
        "user-123",
        "upload-123",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `kardgosdk.OrganizationID` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**uploadID:** `string` — The upload ID identifying the upload session to update
    
</dd>
</dl>

<dl>
<dd>

**request:** `*users.UpdateUploadRequestObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

