# Reference
## Auth
<details><summary><code>client.Auth.GetToken(request) -> *kard.TokenResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &kard.GetTokenRequest{
        ClientId: "client_id",
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

**clientId:** `string` 
    
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
<details><summary><code>client.Files.GetMetadata(OrganizationId) -> *kard.GetFilesMetadataResponse</code></summary>
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
request := &kard.GetFilesMetadataRequest{
        PageSize: kard.Int(
            5,
        ),
        FilterDateFrom: kard.String(
            "2025-02-20T21:56:23Z",
        ),
        FilterDateTo: kard.String(
            "2025-03-20T21:56:23Z",
        ),
        FilterFileType: kard.FileTypeEarnedRewardApprovedDailyReconciliationFile.Ptr(),
        Sort: []*kard.FilesMetadataSortOptions{
            kard.FilesMetadataSortOptionsSentDateDesc.Ptr(),
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

**organizationId:** `kard.OrganizationId` 
    
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

**filterFileType:** `*kard.FileType` — The document file type.
    
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

**sort:** `*kard.FilesMetadataSortOptions` — If provided, response will be sorted by the specified fields. Defaults to descending sentDate, equivalent to "-sentDate"
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Subscriptions
<details><summary><code>client.Notifications.Subscriptions.Get(OrganizationId) -> *notifications.SubscriptionsResponseObject</code></summary>
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**filterEventName:** `*kard.NotificationType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Notifications.Subscriptions.Create(OrganizationId, request) -> *notifications.CreateSubscriptionsResponseObject</code></summary>
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
                        EventName: kard.NotificationTypeEarnedRewardApproved,
                        WebhookUrl: "https://webhookUrl.com/post",
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

**organizationId:** `kard.OrganizationId` 
    
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

<details><summary><code>client.Notifications.Subscriptions.Update(OrganizationId, SubscriptionId, request) -> *notifications.UpdateSubscriptionsResponseObject</code></summary>
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
                    EventName: kard.NotificationTypeEarnedRewardApproved.Ptr(),
                    WebhookUrl: kard.String(
                        "https://webhookUrl.com/post",
                    ),
                    Enabled: kard.Bool(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**subscriptionId:** `kard.SubscriptionId` 
    
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
<details><summary><code>client.Organizations.Get() -> *kard.ExternalOrganizationResponse</code></summary>
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
<details><summary><code>client.Organizations.Children.List(OrganizationId) -> *organizations.ChildOrganizationListResponse</code></summary>
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

**organizationId:** `string` — Unique identifier of the parent organization
    
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

<details><summary><code>client.Organizations.Children.Create(OrganizationId, request) -> *organizations.ChildOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a child organization by cloning the parent and overriding specified fields. An 8-digit numeric ID is generated automatically. The name is required, must contain at least one letter, and may contain only letters and spaces.
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

**organizationId:** `string` — Unique identifier of the parent organization
    
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

<details><summary><code>client.Organizations.Children.Get(OrganizationId, ChildId) -> *organizations.ChildOrganizationResponse</code></summary>
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

**organizationId:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**childId:** `string` — Unique identifier of the child organization
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Children.Update(OrganizationId, ChildId, request) -> *organizations.ChildOrganizationResponse</code></summary>
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

**organizationId:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**childId:** `string` — Unique identifier of the child organization
    
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

<details><summary><code>client.Organizations.Children.Delete(OrganizationId, ChildId) -> *kard.DeleteResourceResponse</code></summary>
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

**organizationId:** `string` — Unique identifier of the parent organization
    
</dd>
</dl>

<dl>
<dd>

**childId:** `string` — Unique identifier of the child organization
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ContentStrategies
<details><summary><code>client.Organizations.ContentStrategies.Create(OrganizationId, request) -> *organizations.ContentStrategyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a content strategy for the organization. The strategy name must be unique within the organization.
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
request := &organizations.CreateContentStrategyRequestBody{
        Data: &organizations.CreateContentStrategyRequestData{
            Attributes: &organizations.CreateContentStrategyAttributes{
                Name: "Featured Travel",
                Sort: organizations.ContentStrategySortHighestCashback.Ptr(),
                Categories: []kard.CategoryOption{
                    kard.CategoryOptionTravel,
                },
                CategoryExclusions: []kard.CategoryOption{
                    kard.CategoryOptionGas,
                },
                MerchantExclusions: []string{
                    "merchant-abc",
                },
            },
        },
    }
client.Organizations.ContentStrategies.Create(
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**request:** `*organizations.CreateContentStrategyRequestBody` — Content strategy data for creation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ContentStrategies.List(OrganizationId) -> *organizations.ContentStrategyListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List content strategies belonging to the authenticated organization
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
request := &organizations.ListContentStrategiesRequest{}
client.Organizations.ContentStrategies.List(
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**filterName:** `*string` — Filter by exact content strategy name (unique within an organization)
    
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

<details><summary><code>client.Organizations.ContentStrategies.Get(OrganizationId, ContentStrategyId) -> *organizations.ContentStrategyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific content strategy
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
client.Organizations.ContentStrategies.Get(
        context.TODO(),
        "organizationId",
        "contentStrategyId",
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**contentStrategyId:** `string` — Unique identifier of the content strategy (UUID v7)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ContentStrategies.Update(OrganizationId, ContentStrategyId, request) -> *organizations.ContentStrategyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace a content strategy. All fields must be provided; any omitted attribute is treated as cleared.
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
request := &organizations.UpdateContentStrategyRequestBody{
        Data: &organizations.UpdateContentStrategyRequestData{
            Attributes: &organizations.UpdateContentStrategyAttributes{
                Name: "name",
                Categories: []kard.CategoryOption{
                    kard.CategoryOptionArtsEntertainment,
                    kard.CategoryOptionArtsEntertainment,
                },
                CategoryExclusions: []kard.CategoryOption{
                    kard.CategoryOptionArtsEntertainment,
                    kard.CategoryOptionArtsEntertainment,
                },
                MerchantExclusions: []string{
                    "merchantExclusions",
                    "merchantExclusions",
                },
            },
        },
    }
client.Organizations.ContentStrategies.Update(
        context.TODO(),
        "organizationId",
        "contentStrategyId",
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**contentStrategyId:** `string` — Unique identifier of the content strategy (UUID v7)
    
</dd>
</dl>

<dl>
<dd>

**request:** `*organizations.UpdateContentStrategyRequestBody` — Content strategy data for update
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ContentStrategies.Delete(OrganizationId, ContentStrategyId) -> *kard.DeleteResourceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a content strategy. Returns 409 if the strategy is still referenced by another resource.
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
client.Organizations.ContentStrategies.Delete(
        context.TODO(),
        "organizationId",
        "contentStrategyId",
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**contentStrategyId:** `string` — Unique identifier of the content strategy (UUID v7)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Placements
<details><summary><code>client.Organizations.Placements.Create(OrganizationId, request) -> *organizations.PlacementFormatUnion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a placement for the organization. Use type "placement" for standard placements (requires name and availableSlots), "placementPushNotification" for push-notification placements (requires name and cadence; availableSlots is automatically set to 1), "placementEmail" for email placements (requires name, cadence, and availableSlots), "placementBatchActivation" for batch-activation placements (requires name, refreshInterval, and slots), or "placementGroup" for group placements (requires name and slots).
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
            Placement: &organizations.CreateStandardPlacementData{
                Attributes: &organizations.CreateStandardAttributes{
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

**organizationId:** `string` — Unique identifier of the organization
    
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

<details><summary><code>client.Organizations.Placements.List(OrganizationId) -> *organizations.PlacementListResponse</code></summary>
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**filterType:** `*organizations.PlacementTypeFilter` — Filter by placement type (placement, placementPushNotification, placementEmail, placementBatchActivation, or placementGroup)
    
</dd>
</dl>

<dl>
<dd>

**filterName:** `*string` — Filter by exact placement name (unique within an organization per type)
    
</dd>
</dl>

<dl>
<dd>

**filterContentStrategyId:** `*string` — Filter by the ID of the content strategy linked to the placement
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — CSV list of related resources to embed in the `included` array. Supported paths: `contentStrategy` (the direct content strategy of a non-batch placement), `slots` (the slot resources of a batch-activation or group placement), `slots.placement` (and the placement each slot references), and `slots.placement.contentStrategy` (and the content strategy of each referenced placement). Dotted paths implicitly include all intermediate resources.
    
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

<details><summary><code>client.Organizations.Placements.Get(OrganizationId, PlacementId) -> *organizations.PlacementResource</code></summary>
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
request := &organizations.GetPlacementRequest{}
client.Organizations.Placements.Get(
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**placementId:** `string` — Unique identifier of the placement (UUID v7)
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — CSV list of related resources to embed in the `included` array. Supported paths: `contentStrategy` (the direct content strategy of a non-batch placement), `slots` (the slot resources of a batch-activation or group placement), `slots.placement` (and the placement each slot references), and `slots.placement.contentStrategy` (and the content strategy of each referenced placement). Dotted paths implicitly include all intermediate resources.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Placements.Update(OrganizationId, PlacementId, request) -> *organizations.PlacementFormatUnion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace a placement. All fields must be provided. Use type "placement", "placementPushNotification", "placementEmail", "placementBatchActivation", or "placementGroup" to set the placement kind. If the type is "placementPushNotification", availableSlots is automatically set to 1.
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
            Placement: &organizations.UpdateStandardPlacementData{
                Attributes: &organizations.UpdateStandardAttributes{
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**placementId:** `string` — Unique identifier of the placement (UUID v7)
    
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

<details><summary><code>client.Organizations.Placements.Delete(OrganizationId, PlacementId) -> *kard.DeleteResourceResponse</code></summary>
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

**organizationId:** `string` — Unique identifier of the organization
    
</dd>
</dl>

<dl>
<dd>

**placementId:** `string` — Unique identifier of the placement (UUID v7)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ping
<details><summary><code>client.Ping.Ping() -> *kard.PingResponseObject</code></summary>
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
<details><summary><code>client.Transactions.Create(OrganizationId, request) -> *kard.TransactionsResponse</code></summary>
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
request := &kard.TransactionsRequestBody{
        Data: []*kard.Transactions{
            &kard.Transactions{
                Transaction: &kard.TransactionsRequest{
                    Id: "309rjfoincor3icno3rind093cdow3jciwjdwcm",
                    Attributes: &kard.TransactionsAttributes{
                        UserId: "6FHt5b6Fnp0qdomMEy5AN6PXcSJIeX69",
                        Status: kard.TransactionStatusApproved,
                        Amount: 1000,
                        Subtotal: kard.Int(
                            800,
                        ),
                        Currency: "USD",
                        Direction: kard.DirectionTypeDebit,
                        PaymentType: kard.TransactionPaymentTypeCard,
                        Description: "ADVANCEAUTO",
                        Description2: kard.String(
                            "ADVANCEAUTO",
                        ),
                        Mcc: kard.String(
                            "1234",
                        ),
                        CardBin: "123456",
                        CardLastFour: "4321",
                        AuthorizationDate: kard.Time(
                            kard.MustParseDateTime(
                                "2021-07-02T17:47:06Z",
                            ),
                        ),
                        Merchant: &kard.Merchant{
                            Id: kard.String(
                                "12345678901234567",
                            ),
                            Name: "ADVANCEAUTO",
                            AddrStreet: kard.String(
                                "125 Main St",
                            ),
                            AddrCity: kard.String(
                                "Philadelphia",
                            ),
                            AddrState: kard.StatesPa.Ptr(),
                            AddrZipcode: kard.String(
                                "19147",
                            ),
                            AddrCountry: kard.String(
                                "United States",
                            ),
                            Latitude: kard.String(
                                "37.9419429",
                            ),
                            Longitude: kard.String(
                                "-73.1446869",
                            ),
                            StoreId: kard.String(
                                "12345",
                            ),
                        },
                        AuthorizationCode: kard.String(
                            "123456",
                        ),
                        RetrievalReferenceNumber: kard.String(
                            "100804333919",
                        ),
                        AcquirerReferenceNumber: kard.String(
                            "1234567890123456789012345678",
                        ),
                        SystemTraceAuditNumber: kard.String(
                            "333828",
                        ),
                        TransactionId: "2467de37-cbdc-416d-a359-75de87bfffb0",
                        CardProductId: kard.String(
                            "1234567890123456789012345678",
                        ),
                        ProcessorMids: &kard.ProcessorMid{
                            Visa: &kard.VisaMid{
                                Mids: &kard.VisaMidDetails{
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kard.TransactionsRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.CreateFraudMarkers(OrganizationId, request) -> *kard.FraudulentTransactionObject</code></summary>
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
request := &kard.FraudulentTransactionRequestBody{
        Data: []*kard.FraudulentTransactionData{
            &kard.FraudulentTransactionData{
                Id: "myTxnId12345",
                Type: "fraudulentTransaction",
                Attributes: &kard.FraudulentTransactionAttributes{
                    UserId: "userId123",
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kard.FraudulentTransactionRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.CreateAudits(OrganizationId, UserId, request) -> *kard.CreateAuditResponseBody</code></summary>
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
request := &kard.CreateAuditRequestBody{
        Data: []*kard.CreateAuditRequestDataUnion{
            &kard.CreateAuditRequestDataUnion{
                Audit: &kard.AuditRequestData{
                    Attributes: &kard.AuditAttributes{
                        AuditCode: 8001,
                        MerchantName: "Caribbean Goodness",
                        AuditDescription: "duplicate transaction",
                        TransactionId: "issuerTransaction123",
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kard.CreateAuditRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.CreateBulkTransactionsUploadUrl(OrganizationId, request) -> *kard.CreateFileUploadUrlResponse</code></summary>
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
Supports both `incomingTransactionsFile` for daily transaction ingestion and `historicalTransactionsFile` for historical transaction ingestion. See the [Historical Transaction Uploads](/2024-10-01/api/integration-guides/historical-transaction-uploads) integration guide for details on the historical flow.
<b>Required scopes:</b> `files:write`
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
request := &kard.CreateFileUploadRequestBody{
        Data: []*kard.CreateFileUploadData{
            &kard.CreateFileUploadData{
                Type: kard.FileUploadTypeIncomingTransactionsFile,
                Attributes: &kard.CreateFileUploadAttributes{
                    Filename: "transaction_12345.jsonl",
                },
            },
            &kard.CreateFileUploadData{
                Type: kard.FileUploadTypeIncomingTransactionsFile,
                Attributes: &kard.CreateFileUploadAttributes{
                    Filename: "transaction_67890.jsonl",
                },
            },
        },
    }
client.Transactions.CreateBulkTransactionsUploadUrl(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kard.CreateFileUploadRequestBody` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.GetEarnedRewards(OrganizationId, UserId) -> *kard.GetEarnedRewardsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve rewarded transaction history for a specific user. By default this returns only SETTLED transactions within the last 12 months regardless of payment status. Pass `filter[paidInFullOnly]=true` to restrict the response to matched transactions that have been paid in full to the issuer (`paidToIssuer` is `PAID_IN_FULL`).
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
request := &kard.GetEarnedRewardsRequest{
        PageSize: kard.Int(
            10,
        ),
        FilterStatus: kard.RewardedTransactionStatusApproved.Ptr(),
        Include: kard.String(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` — The ID of the user as defined on the issuers system
    
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

**filterStatus:** `*kard.RewardedTransactionStatus` — Filter by transaction status. Supported values are `APPROVED` and `SETTLED`. Defaults to `SETTLED` when omitted. When `APPROVED` is specified, only approved transactions that do not yet have a corresponding settled transaction are returned.
    
</dd>
</dl>

<dl>
<dd>

**filterPaidInFullOnly:** `*bool` — When `true`, only return transactions that have been paid in full to the issuer (`paidToIssuer` is `PAID_IN_FULL`). By default (`false`), any matched transaction is returned regardless of payment status. This also controls whether unpaid transactions contribute to `lifetimeRewardsInCents`. Has no effect on `APPROVED` transactions, which are always returned when requested.
    
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
<details><summary><code>client.Users.Create(OrganizationId, request) -> *kard.CreateUsersObject</code></summary>
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
request := &kard.CreateUsersObject{
        Data: []*kard.UserRequestDataUnion{
            &kard.UserRequestDataUnion{
                User: &kard.UserRequestData{
                    Id: "1234567890",
                    Attributes: &kard.UserRequestAttributes{
                        ZipCode: kard.String(
                            "11238",
                        ),
                        EnrolledRewards: []kard.EnrolledRewardsType{
                            kard.EnrolledRewardsTypeCardlinked,
                        },
                        Email: kard.String(
                            "user@example.com",
                        ),
                        HashedEmail: kard.String(
                            "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3e2d8a5b76e45a1d4c4e2e3a1",
                        ),
                        PhoneNumber: kard.String(
                            "+14155552671",
                        ),
                        BirthYear: kard.String(
                            "1990",
                        ),
                        HistoricalTransactionsSent: kard.Bool(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kard.CreateUsersObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Update(OrganizationId, UserId, request) -> *kard.UserResponseObject</code></summary>
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
request := &kard.UpdateUserObject{
        Data: &kard.UpdateUserRequestDataUnion{
            User: &kard.UpdateUserRequestData{
                Id: "1234567890",
                Attributes: &kard.UpdateUserRequestAttributes{
                    ZipCode: kard.String(
                        "11238",
                    ),
                    EnrolledRewards: []kard.EnrolledRewardsType{
                        kard.EnrolledRewardsTypeCardlinked,
                    },
                    Email: kard.String(
                        "user@example.com",
                    ),
                    HashedEmail: kard.String(
                        "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3e2d8a5b76e45a1d4c4e2e3a1",
                    ),
                    PhoneNumber: kard.String(
                        "+14155552671",
                    ),
                    BirthYear: kard.String(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*kard.UpdateUserObject` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Delete(OrganizationId, UserId) -> *kard.DeleteUserResponseObject</code></summary>
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Get(OrganizationId, UserId) -> *kard.UserResponseObject</code></summary>
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Attributions
<details><summary><code>client.Users.Attributions.Create(OrganizationId, UserId, request) -> *users.CreateAttributionResponse</code></summary>
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
                        EntityId: "60e4ba1da31c5a22a144c075",
                        EventCode: users.EventCodeView,
                        Medium: users.OfferMediumSearch,
                        EventDate: kard.MustParseDateTime(
                            "2025-01-01T00:00:00Z",
                        ),
                    },
                },
            },
            &users.CreateAttributionRequestUnion{
                OfferAttribution: &users.OfferAttributionRequest{
                    Attributes: &users.OfferAttributionAttributes{
                        EntityId: "60e4ba1da31c5a22a144c077",
                        EventCode: users.EventCodeImpression,
                        Medium: users.OfferMediumEmail,
                        EventDate: kard.MustParseDateTime(
                            "2025-01-01T00:00:00Z",
                        ),
                    },
                },
            },
            &users.CreateAttributionRequestUnion{
                NotificationAttribution: &users.NotificationAttributionRequest{
                    Attributes: &users.NotificationAttributionAttributes{
                        EntityId: "60e4ba1da31c5a22a144c076",
                        EventCode: users.EventCodeImpression,
                        Medium: users.NotificationMediumPush,
                        EventDate: kard.MustParseDateTime(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
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

<details><summary><code>client.Users.Attributions.Activate(OrganizationId, UserId, OfferId) -> *users.ActivateOfferResponse</code></summary>
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**offerId:** `string` — The unique identifier of the offer being activated
    
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

<details><summary><code>client.Users.Attributions.Boost(OrganizationId, UserId, OfferId) -> *users.BoostOfferResponse</code></summary>
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**offerId:** `string` — The unique identifier of the offer being boosted
    
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

<details><summary><code>client.Users.Attributions.ActivatePlacementSlot(OrganizationId, UserId, PlacementId, SlotId) -> *users.ActivatePlacementSlotResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Record when a user activates a batch-activation placement slot. Writes a slot-level
`placementSlotAttribution` ACTIVATE event and fans out a per-offer
`offerAttribution` ACTIVATE event for every offer resolved by the slot's content
strategy. The slot-level event id and the resolved `offerIds` are returned so the
partner can render the batch immediately without an extra `getBatchesByPlacement`
round-trip.

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
client.Users.Attributions.ActivatePlacementSlot(
        context.TODO(),
        "organization-123",
        "user-123",
        "018f8d6b-1abc-7def-9012-345678901234",
        "slot-a",
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**placementId:** `string` — Unique identifier of the placement (UUID v7)
    
</dd>
</dl>

<dl>
<dd>

**slotId:** `string` — Stable identifier for the slot within the placement
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## WebView
<details><summary><code>client.Users.Auth.GetWebViewToken(OrganizationId, UserId) -> *users.WebViewTokenResponse</code></summary>
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Rewards
<details><summary><code>client.Users.Rewards.Offers(OrganizationId, UserId) -> *users.OffersResponseObject</code></summary>
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
        PageSize: kard.Int(
            1,
        ),
        FilterIsTargeted: kard.Bool(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
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

**filterPurchaseChannel:** `[]*kard.PurchaseChannel` 
    
</dd>
</dl>

<dl>
<dd>

**filterCategory:** `*kard.CategoryOption` 
    
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

<details><summary><code>client.Users.Rewards.PlacementOffers(OrganizationId, UserId, PlacementId) -> *users.OffersResponseObject</code></summary>
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**placementId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**filterSearch:** `*string` — Case-insensitive search string to filter offers by merchant name
    
</dd>
</dl>

<dl>
<dd>

**filterPurchaseChannel:** `[]*kard.PurchaseChannel` 
    
</dd>
</dl>

<dl>
<dd>

**filterCategory:** `*kard.CategoryOption` 
    
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

<details><summary><code>client.Users.Rewards.PlacementBatches(OrganizationId, UserId, PlacementId) -> *users.BatchesResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve batches for a batch-activation or group placement. Returns each
slot in slot order with its current offer set, alias, and freshness fields
(`isActive`, `lastActivatedAt`, `expiresAt`). Applies the same per-user
eligibility and per-slot content-strategy filter as Get Offers By
Placement, independently per slot. For a batch-activation placement, a
slot only flips to `isActive: false` when its refresh interval has elapsed
AND its post-eligibility `offers[]` is non-empty; otherwise the slot is
still returned and stays active so the partner UI does not promote
"refresh" with nothing to show. For a group placement, slots are always
active and each slot returns its offers regardless of activation state,
hiding only offers that require activation (`requiredInBatch`) and have
no activation record.<br/>
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
request := &users.GetBatchesByPlacementRequest{}
client.Users.Rewards.PlacementBatches(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**placementId:** `string` 
    
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

<details><summary><code>client.Users.Rewards.PlacementContent(OrganizationId, UserId, PlacementId) -> *users.PlacementContentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the content for a placement. The placement type is resolved
server-side so callers no longer pick an endpoint by placement type.
Returns a JSON:API document whose `data` resources are self-describing
by `type`: a standard placement returns `standardOffer` resources (the
same payload as Get Offers By Placement — with `links`, optional
`included` categories, and `meta`); a batch-activation or group
placement returns `placementBatch` slot resources (the same payload as
Get Batches By Placement). Distinguish the two by each resource's
`type`. Email and push-notification placements are not servable through
this endpoint and respond with a `400`.<br/>
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
request := &users.GetPlacementContentRequest{
        Include: []*string{
            kard.String(
                "categories",
            ),
        },
    }
client.Users.Rewards.PlacementContent(
        context.TODO(),
        "organization-123",
        "user-123",
        "placement-homepage-banner",
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
</dd>
</dl>

<dl>
<dd>

**placementId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` — CSV list of included resources in the response (e.g "categories"). Allowed value is `categories`. Only applies to standard placements (those returning `standardOffer` resources).
    
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

<details><summary><code>client.Users.Rewards.Locations(OrganizationId, UserId) -> *users.LocationsResponseObject</code></summary>
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
        PageSize: kard.Int(
            1,
        ),
        FilterLatitude: kard.Float64(
            39.9419429,
        ),
        FilterLongitude: kard.Float64(
            -75.1446869,
        ),
        FilterRadius: kard.Int(
            10,
        ),
        Include: []*string{
            kard.String(
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `kard.UserId` 
    
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

**filterState:** `*kard.State` 
    
</dd>
</dl>

<dl>
<dd>

**filterCategory:** `*kard.CategoryOption` 
    
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
<details><summary><code>client.Users.Uploads.Create(OrganizationId, UserId, request) -> *users.CreateUploadResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

<b>Deprecated.</b> This endpoint is deprecated in favor of the [Create Bulk Transactions Upload URL](/2024-10-01/api/transactions/create-bulk-transactions-upload-url) endpoint. New integrations should use the bulk flow outlined in the [Historical Transaction Uploads](/2024-10-01/api/integration-guides/historical-transaction-uploads) integration guide.

Call this endpoint to create an upload session and retrieve an upload ID. Using the upload ID in the [Add Upload Part](/2024-10-01/api/transactions/uploads/create-part) endpoint, historical transactions can be sent in batches for further processing.
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
                Attributes: &kard.EmptyObject{},
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` — The ID of the user as defined on the issuers system
    
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

<details><summary><code>client.Users.Uploads.CreatePart(OrganizationId, UserId, UploadId, request) -> *users.CreateUploadPartResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

<b>Deprecated.</b> This endpoint is deprecated in favor of the [Create Bulk Transactions Upload URL](/2024-10-01/api/transactions/create-bulk-transactions-upload-url) endpoint. New integrations should use the bulk flow outlined in the [Historical Transaction Uploads](/2024-10-01/api/integration-guides/historical-transaction-uploads) integration guide.

Call this endpoint using the upload ID provided in the [Create Upload](/2024-10-01/api/transactions/uploads/create) endpoint to add parts to your upload. Currently, this endpoint supports adding historical transactions.
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
                HistoricalTransaction: &kard.TransactionsRequest{
                    Id: "309rjfoincor3icno3rind093cdow3jciwjdwcm",
                    Attributes: &kard.TransactionsAttributes{
                        UserId: "6FHt5b6Fnp0qdomMEy5AN6PXcSJIeX69",
                        Status: kard.TransactionStatusApproved,
                        Amount: 1000,
                        Subtotal: kard.Int(
                            800,
                        ),
                        Currency: "USD",
                        Direction: kard.DirectionTypeDebit,
                        PaymentType: kard.TransactionPaymentTypeCard,
                        Description: "ADVANCEAUTO",
                        Description2: kard.String(
                            "ADVANCEAUTO",
                        ),
                        Mcc: kard.String(
                            "1234",
                        ),
                        CardBin: "123456",
                        CardLastFour: "4321",
                        AuthorizationDate: kard.Time(
                            kard.MustParseDateTime(
                                "2021-07-02T17:47:06Z",
                            ),
                        ),
                        Merchant: &kard.Merchant{
                            Id: kard.String(
                                "12345678901234567",
                            ),
                            Name: "ADVANCEAUTO",
                            AddrStreet: kard.String(
                                "125 Main St",
                            ),
                            AddrCity: kard.String(
                                "Philadelphia",
                            ),
                            AddrState: kard.StatesPa.Ptr(),
                            AddrZipcode: kard.String(
                                "19147",
                            ),
                            AddrCountry: kard.String(
                                "United States",
                            ),
                            Latitude: kard.String(
                                "37.9419429",
                            ),
                            Longitude: kard.String(
                                "-73.1446869",
                            ),
                            StoreId: kard.String(
                                "12345",
                            ),
                        },
                        AuthorizationCode: kard.String(
                            "123456",
                        ),
                        RetrievalReferenceNumber: kard.String(
                            "100804333919",
                        ),
                        AcquirerReferenceNumber: kard.String(
                            "1234567890123456789012345678",
                        ),
                        SystemTraceAuditNumber: kard.String(
                            "333828",
                        ),
                        TransactionId: "2467de37-cbdc-416d-a359-75de87bfffb0",
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**uploadId:** `string` — The upload ID identifying the upload session to add parts
    
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

<details><summary><code>client.Users.Uploads.Update(OrganizationId, UserId, UploadId, request) -> *users.UpdateUploadResponseObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

<b>Deprecated.</b> This endpoint is deprecated in favor of the [Create Bulk Transactions Upload URL](/2024-10-01/api/transactions/create-bulk-transactions-upload-url) endpoint. New integrations should use the bulk flow outlined in the [Historical Transaction Uploads](/2024-10-01/api/integration-guides/historical-transaction-uploads) integration guide.

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
                Id: "7e382223-b9a5-4825-91fb-436c8957a2e7",
                Attributes: &kard.EmptyObject{},
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

**organizationId:** `kard.OrganizationId` 
    
</dd>
</dl>

<dl>
<dd>

**userId:** `string` — The ID of the user as defined on the issuers system
    
</dd>
</dl>

<dl>
<dd>

**uploadId:** `string` — The upload ID identifying the upload session to update
    
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

