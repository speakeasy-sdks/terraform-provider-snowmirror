// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"snowmirror/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"snowmirror/internal/sdk/pkg/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SynchronizationResource{}
var _ resource.ResourceWithImportState = &SynchronizationResource{}

func NewSynchronizationResource() resource.Resource {
	return &SynchronizationResource{}
}

// SynchronizationResource defines the resource implementation.
type SynchronizationResource struct {
	client *sdk.Snowmirror
}

// SynchronizationResourceModel describes the resource data model.
type SynchronizationResourceModel struct {
	Description types.String                    `tfsdk:"description"`
	ID          types.Int64                     `tfsdk:"id"`
	Image       types.String                    `tfsdk:"image"`
	Name        types.String                    `tfsdk:"name"`
	Price       types.Number                    `tfsdk:"price"`
	Sync        *CreateSynchronizationInputSync `tfsdk:"sync"`
	Teaser      types.String                    `tfsdk:"teaser"`
}

func (r *SynchronizationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_synchronization"
}

func (r *SynchronizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Synchronization Resource",

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `Product description of the coffee.`,
			},
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: `Sync ID`,
			},
			"image": schema.StringAttribute{
				Computed:    true,
				Description: `URI for an image of the coffee.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Product name of the coffee.`,
			},
			"price": schema.NumberAttribute{
				Computed:    true,
				Description: `Suggested cost of the coffee.`,
			},
			"sync": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Optional: true,
					},
					"allow_inherited_columns": schema.BoolAttribute{
						Optional:    true,
						Description: `SnowMirror checks if columns exist in ServiceNow. If this flag is set to true,`,
					},
					"auto_schema_update": schema.BoolAttribute{
						Optional: true,
					},
					"columns": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"columns_to_exclude": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"delete_strategy": schema.StringAttribute{
						Optional: true,
					},
					"encoded_query": schema.StringAttribute{
						Optional: true,
					},
					"full_load_scheduler": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"begin_date": schema.StringAttribute{
								Optional: true,
							},
							"execution_type": schema.StringAttribute{
								Optional: true,
							},
							"type": schema.StringAttribute{
								Optional: true,
							},
						},
					},
					"mirror_table": schema.StringAttribute{
						Required:    true,
						Description: `Name of the table in mirror database where the data will be migrated.`,
					},
					"name": schema.StringAttribute{
						Required:    true,
						Description: `Display name of the synchronization.`,
					},
					"reference_field_type": schema.StringAttribute{
						Optional:    true,
						Description: `Defines how to synchronize reference field types.`,
					},
					"run_immediately": schema.BoolAttribute{
						Optional: true,
					},
					"scheduler": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"begin_date": schema.StringAttribute{
								Optional: true,
							},
							"type": schema.StringAttribute{
								Optional: true,
							},
						},
					},
					"scheduler_priority": schema.StringAttribute{
						Optional: true,
					},
					"table": schema.StringAttribute{
						Optional:    true,
						Description: `Name of the table in ServiceNow.`,
					},
					"view": schema.StringAttribute{
						Optional:    true,
						Description: `Name of the view in ServiceNow.`,
					},
				},
			},
			"teaser": schema.StringAttribute{
				Computed:    true,
				Description: `Fun tagline for the coffee.`,
			},
		},
	}
}

func (r *SynchronizationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Snowmirror)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Snowmirror, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SynchronizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SynchronizationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToCreateSDKType()
	res, err := r.client.Synchronization.CreateSynchronization(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Synchronization == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.Synchronization)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SynchronizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SynchronizationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueInt64()
	request := operations.GetSynchronizationRequest{
		ID: id,
	}
	res, err := r.client.Synchronization.GetSynchronization(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Synchronization == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Synchronization)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SynchronizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SynchronizationResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	createSynchronizationInput := *data.ToUpdateSDKType()
	id := data.ID.ValueInt64()
	request := operations.UpdateSynchronizationRequest{
		CreateSynchronizationInput: createSynchronizationInput,
		ID:                         id,
	}
	res, err := r.client.Synchronization.UpdateSynchronization(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Synchronization == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.Synchronization)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SynchronizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SynchronizationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueInt64()
	request := operations.DeleteSynchronizationRequest{
		ID: id,
	}
	res, err := r.client.Synchronization.DeleteSynchronization(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SynchronizationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}