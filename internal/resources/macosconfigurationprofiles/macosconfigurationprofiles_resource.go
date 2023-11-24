// macosconfigurationprofiles_resource.go
package macosconfigurationprofiles

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/terraform-provider-jamfpro/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceJamfProMacOSConfigurationProfiles defines the schema and CRUD operations for managing Jamf Pro Departments in Terraform.
func ResourceJamfProMacOSConfigurationProfiles() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceJamfProMacOSConfigurationProfilesCreate,
		ReadContext:   ResourceJamfProMacOSConfigurationProfilesRead,
		UpdateContext: ResourceJamfProMacOSConfigurationProfilesUpdate,
		DeleteContext: ResourceJamfProMacOSConfigurationProfilesDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			// GeneralConfig fields
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the macOS configuration profile.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the configuration profile.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the configuration profile.",
			},
			"site": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Jamf Pro Site information for the assigned configuration profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "-1", // Set default value as string "-1"
							Description: "Jamf Pro Site ID. Value defaults to -1 aka not used.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "None", // Set default value as "None"
							Description: "Jamf Pro Site Name. Value defaults to 'None' aka not used",
						},
					},
				},
			},
			"category": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Configuration Profile Category information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "-1", // Set default value as string "-1"
							Description: "Category ID. Value defaults to -1 aka not used.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "No category assigned", // Set default value as "No category assigned"
							Description: "Category Name for assigned configuration profile. Value defaults to 'No category assigned' aka not used",
						},
					},
				},
			},
			"distribution_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The distribution methods for the macOS configuration profile, can be either Install Automatically or Make Available in Self Service.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					value := val.(string)
					allowedValues := []string{"Install Automatically", "Make Available in Self Service"}
					for _, v := range allowedValues {
						if value == v {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be one of [%s], got: %s", key, strings.Join(allowedValues, ", "), value))
					return
				},
			},
			"user_removable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Define whether the macOS configuration profile is removeable by the end user using jamf self service.",
			},
			"level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The level defines what level the MDM profile is deployed at. It can either be device wide using 'System'', or for an individual user using 'User'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					value := val.(string)
					allowedValues := []string{"System", "User"}
					for _, v := range allowedValues {
						if value == v {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be either 'System' or 'User', got: %s", key, value))
					return
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The uuid of the macos configuration profile",
			},
			"redeploy_on_update": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Newly Assigned",
				Description: "",
			},
			"payloads": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration profile payload in xml and delivered as a plist to the macOS device by Jamf Pro.",
			},
			// ScopeConfig fields
			"scope": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Scope configuration for the profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_computers": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, applies the profile to all computers.",
						},
						"all_jss_users": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, applies the profile to all JSS users.",
						},
						"computers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"computer": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The unique identifier of the scoped computer.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the scoped computer.",
												},
												"udid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "UDID of the scoped computer.",
												},
											},
										},
									},
								},
							},
						},
						"buildings": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"building": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The unique identifier of the scoped building.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the scoped building.",
												},
											},
										},
									},
								},
							},
						},
						"departments": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"department": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The unique identifier of the scoped department.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the scoped department.",
												},
											},
										},
									},
								},
							},
						},
						"computer_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"computer_group": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The unique identifier of the scoped computer group.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the computer scoped group.",
												},
											},
										},
									},
								},
							},
						},
						"jss_users": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"jss_user": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The unique identifier of the scoped JSS user.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the scoped JSS user.",
												},
											},
										},
									},
								},
							},
						},
						"jss_user_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The unique identifier of the scoped JSS user group.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the scoped JSS user group.",
									},
								},
							},
						},
						"limitations": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the user.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the user.",
															},
														},
													},
												},
											},
										},
									},
									"user_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_group": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the scoped user group.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the scoped user group.",
															},
														},
													},
												},
											},
										},
									},
									"network_segments": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"network_segment": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the scoped network segment.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the scoped network segment.",
															},
															"uid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "UID of the scoped network segment.",
															},
														},
													},
												},
											},
										},
									},
									"ibeacons": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ibeacon": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the scoped iBeacon.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the scoped iBeacon.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"exclusions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"computers": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"computer": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the computer.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the computer.",
															},
															"udid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "UDID of the computer.",
															},
														},
													},
												},
											},
										},
									},
									"buildings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"building": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the building.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the building.",
															},
														},
													},
												},
											},
										},
									},
									"departments": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"department": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the department.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the department.",
															},
														},
													},
												},
											},
										},
									},
									"computer_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"computer_group": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the computer group.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the computer group.",
															},
														},
													},
												},
											},
										},
									},
									"users": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the user.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the user.",
															},
														},
													},
												},
											},
										},
									},
									"user_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_group": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the user group.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the user group.",
															},
														},
													},
												},
											},
										},
									},
									"network_segments": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"network_segment": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the network segment.",
															},
															"uid": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "UID of the network segment.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the network segment.",
															},
														},
													},
												},
											},
										},
									},
									"ibeacons": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ibeacon": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the iBeacon.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the iBeacon.",
															},
														},
													},
												},
											},
										},
									},
									"jss_users": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"jss_user": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the JSS user.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the JSS user.",
															},
														},
													},
												},
											},
										},
									},
									"jss_user_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"jss_user_group": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The unique identifier of the JSS user group.",
															},
															"name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the JSS user group.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			// SelfServiceConfig fields
			"self_service": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Self Service configuration for the profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"install_button_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"self_service_description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"force_users_to_view_description": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"self_service_icon": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"data": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"feature_on_main_page": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"self_service_categories": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of categories under self service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The category information for the self service.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The unique identifier of the category.",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the category.",
												},
												"display_in": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to display in self-service.",
												},
												"feature_in": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to feature in self-service.",
												},
											},
										},
									},
								},
							},
						},
						"notification": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"notification_subject": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"notification_message": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// constructJamfProMacOSConfigurationProfile constructs a ResponseMacOSConfigurationProfile object from the provided schema data.
// It captures each attribute from the schema and returns the constructed ResponseMacOSConfigurationProfile object.
func constructJamfProMacOSConfigurationProfile(d *schema.ResourceData) *jamfpro.ResponseMacOSConfigurationProfiles {
	// Construct General section
	// Initialize an empty General struct
	general := jamfpro.MacOSConfigurationProfilesDataSubsetGeneral{}

	// Safely retrieve and set the "name" field if it exists
	if v, ok := d.GetOk("name"); ok {
		general.Name = v.(string)
	}

	// Safely retrieve and set the "description" field if it exists
	if v, ok := d.GetOk("description"); ok {
		general.Description = v.(string)
	}

	// Safely retrieve and set the "distribution_method" field if it exists
	if v, ok := d.GetOk("distribution_method"); ok {
		general.DistributionMethod = v.(string)
	}

	// Safely retrieve and set the "user_removable" field if it exists
	if v, ok := d.GetOk("user_removable"); ok {
		general.UserRemovable = v.(bool)
	}

	// Safely retrieve and set the "level" field if it exists
	if v, ok := d.GetOk("level"); ok {
		general.Level = v.(string)
	}

	// Safely retrieve and set the "uuid" field if it exists
	if v, ok := d.GetOk("uuid"); ok {
		general.UUID = v.(string)
	}

	// Safely retrieve and set the "redeploy_on_update" field if it exists
	if v, ok := d.GetOk("redeploy_on_update"); ok {
		general.RedeployOnUpdate = v.(string)
	}

	// Safely retrieve and set the "payloads" field if it exists
	if v, ok := d.GetOk("payloads"); ok {
		general.Payloads = v.(string)
	}

	// Safely construct the Site field from the Terraform schema
	if siteList, ok := d.GetOk("site"); ok && len(siteList.([]interface{})) > 0 {
		siteMap := siteList.([]interface{})[0].(map[string]interface{})

		var siteID int
		var siteName string

		// Convert the site ID from string to int and check for presence
		if idStr, ok := siteMap["id"].(string); ok {
			siteID, _ = strconv.Atoi(idStr) // Convert string to int, ignoring error as default value is a valid int
		}

		// Check for presence of site name
		if name, ok := siteMap["name"].(string); ok {
			siteName = name
		}

		general.Site = jamfpro.MacOSConfigurationProfilesDataSubsetSite{
			ID:   siteID,
			Name: siteName,
		}
	}

	// Safely construct the Category field from the Terraform schema
	if categoryList, ok := d.GetOk("category"); ok && len(categoryList.([]interface{})) > 0 {
		categoryMap := categoryList.([]interface{})[0].(map[string]interface{})

		var categoryID int
		var categoryName string

		// Convert the category ID from string to int and check for presence
		if idStr, ok := categoryMap["id"].(string); ok {
			categoryID, _ = strconv.Atoi(idStr) // Convert string to int, ignoring error as default value is a valid int
		}

		// Check for presence of category name
		if name, ok := categoryMap["name"].(string); ok {
			categoryName = name
		}

		general.Category = jamfpro.MacOSConfigurationProfilesDataSubsetCategory{
			ID:   categoryID,
			Name: categoryName,
		}
	}

	// Construct Scope section
	scopeData := d.Get("scope").([]interface{})
	var scope jamfpro.MacOSConfigurationProfilesDataSubsetScope
	if len(scopeData) > 0 {
		scopeMap := scopeData[0].(map[string]interface{})

		// Safely set the AllComputers and AllJSSUsers fields
		if v, ok := scopeMap["all_computers"].(bool); ok {
			scope.AllComputers = v
		}
		if v, ok := scopeMap["all_jss_users"].(bool); ok {
			scope.AllJSSUsers = v
		}

		// Safely construct and append Computers
		if computers, ok := scopeMap["computers"].([]interface{}); ok {
			for _, comp := range computers {
				compMap := comp.(map[string]interface{})
				var compID int
				var compName, compUDID string

				if id, ok := compMap["id"].(int); ok {
					compID = id
				}
				if name, ok := compMap["name"].(string); ok {
					compName = name
				}
				if udid, ok := compMap["udid"].(string); ok {
					compUDID = udid
				}

				computer := jamfpro.MacOSConfigurationProfilesDataSubsetComputer{
					Computer: jamfpro.MacOSConfigurationProfilesDataSubsetComputerItem{
						ID:   compID,
						Name: compName,
						UDID: compUDID,
					},
				}
				scope.Computers = append(scope.Computers, computer)
			}
		}

		// Similar pattern for Buildings, Departments, ComputerGroups, JSSUsers, JSSUserGroups
		if buildings, ok := scopeMap["buildings"].([]interface{}); ok {
			for _, bld := range buildings {
				bldMap := bld.(map[string]interface{})
				var bldID int
				var bldName string

				if id, ok := bldMap["id"].(int); ok {
					bldID = id
				}
				if name, ok := bldMap["name"].(string); ok {
					bldName = name
				}

				building := jamfpro.MacOSConfigurationProfilesDataSubsetBuilding{
					Building: jamfpro.MacOSConfigurationProfilesDataSubsetBuildingItem{
						ID:   bldID,
						Name: bldName,
					},
				}
				scope.Buildings = append(scope.Buildings, building)
			}
		}

		// Safely construct and append Departments
		if departments, ok := scopeMap["departments"].([]interface{}); ok {
			for _, dept := range departments {
				deptMap := dept.(map[string]interface{})
				var deptID int
				var deptName string

				if id, ok := deptMap["id"].(int); ok {
					deptID = id
				}
				if name, ok := deptMap["name"].(string); ok {
					deptName = name
				}

				department := jamfpro.MacOSConfigurationProfilesDataSubsetDepartment{
					Department: jamfpro.MacOSConfigurationProfilesDataSubsetDepartmentItem{
						ID:   deptID,
						Name: deptName,
					},
				}
				scope.Departments = append(scope.Departments, department)
			}
		}

		// Safely construct and append ComputerGroups
		if computerGroups, ok := scopeMap["computer_groups"].([]interface{}); ok {
			for _, grp := range computerGroups {
				grpMap := grp.(map[string]interface{})
				var grpID int
				var grpName string

				if id, ok := grpMap["id"].(int); ok {
					grpID = id
				}
				if name, ok := grpMap["name"].(string); ok {
					grpName = name
				}

				computerGroup := jamfpro.MacOSConfigurationProfilesDataSubsetComputerGroup{
					ComputerGroup: jamfpro.MacOSConfigurationProfilesDataSubsetComputerGroupItem{
						ID:   grpID,
						Name: grpName,
					},
				}
				scope.ComputerGroups = append(scope.ComputerGroups, computerGroup)
			}
		}

		// Safely construct and append JSSUsers
		if jssUsers, ok := scopeMap["jss_users"].([]interface{}); ok {
			for _, usr := range jssUsers {
				usrMap := usr.(map[string]interface{})
				var usrID int
				var usrName string

				if id, ok := usrMap["id"].(int); ok {
					usrID = id
				}
				if name, ok := usrMap["name"].(string); ok {
					usrName = name
				}

				jssUser := jamfpro.MacOSConfigurationProfilesDataSubsetJSSUser{
					JSSUser: jamfpro.MacOSConfigurationProfilesDataSubsetJSSUserItem{
						ID:   usrID,
						Name: usrName,
					},
				}
				scope.JSSUsers = append(scope.JSSUsers, jssUser)
			}
		}

		// Safely construct and append JSSUserGroups
		if jssUserGroups, ok := scopeMap["jss_user_groups"].([]interface{}); ok {
			for _, grp := range jssUserGroups {
				grpMap := grp.(map[string]interface{})
				var grpID int
				var grpName string

				if id, ok := grpMap["id"].(int); ok {
					grpID = id
				}
				if name, ok := grpMap["name"].(string); ok {
					grpName = name
				}

				jssUserGroup := jamfpro.MacOSConfigurationProfilesDataSubsetJSSUserGroup{
					JSSUserGroup: jamfpro.MacOSConfigurationProfilesDataSubsetJSSUserGroupItem{
						ID:   grpID,
						Name: grpName,
					},
				}
				scope.JSSUserGroups = append(scope.JSSUserGroups, jssUserGroup)
			}
		}

		// Construct Limitations
		if limitations, ok := scopeMap["limitations"].([]interface{}); ok && len(limitations) > 0 {
			limMap := limitations[0].(map[string]interface{})
			var lim jamfpro.MacOSConfigurationProfilesDataSubsetLimitations

			// Safely construct and append Users in Limitations
			if users, ok := limMap["users"].([]interface{}); ok {
				for _, usr := range users {
					usrMap := usr.(map[string]interface{})
					var usrID int
					var usrName string

					if id, ok := usrMap["id"].(int); ok {
						usrID = id
					}
					if name, ok := usrMap["name"].(string); ok {
						usrName = name
					}

					user := jamfpro.MacOSConfigurationProfilesDataSubsetUser{
						User: jamfpro.MacOSConfigurationProfilesDataSubsetUserItem{
							ID:   usrID,
							Name: usrName,
						},
					}
					lim.Users = append(lim.Users, user)
				}
			}

			// Safely construct and append UserGroups in Limitations
			if userGroups, ok := limMap["user_groups"].([]interface{}); ok {
				for _, grp := range userGroups {
					grpMap := grp.(map[string]interface{})
					var grpID int
					var grpName string

					if id, ok := grpMap["id"].(int); ok {
						grpID = id
					}
					if name, ok := grpMap["name"].(string); ok {
						grpName = name
					}

					userGroup := jamfpro.MacOSConfigurationProfilesDataSubsetUserGroup{
						UserGroup: jamfpro.MacOSConfigurationProfilesDataSubsetUserGroupItem{
							ID:   grpID,
							Name: grpName,
						},
					}
					lim.UserGroups = append(lim.UserGroups, userGroup)
				}
			}

			// Safely construct and append NetworkSegments in Limitations
			if networkSegments, ok := limMap["network_segments"].([]interface{}); ok {
				for _, seg := range networkSegments {
					segMap := seg.(map[string]interface{})
					var segID int
					var segUID, segName string

					if id, ok := segMap["id"].(int); ok {
						segID = id
					}
					if uid, ok := segMap["uid"].(string); ok {
						segUID = uid
					}
					if name, ok := segMap["name"].(string); ok {
						segName = name
					}

					networkSegment := jamfpro.MacOSConfigurationProfilesDataSubsetNetworkSegment{
						NetworkSegment: jamfpro.MacOSConfigurationProfilesDataSubsetNetworkSegmentItem{
							ID:   segID,
							UID:  segUID,
							Name: segName,
						},
					}
					lim.NetworkSegments = append(lim.NetworkSegments, networkSegment)
				}
			}

			// Safely construct and append IBeacons in Limitations
			if ibeacons, ok := limMap["ibeacons"].([]interface{}); ok {
				for _, ibc := range ibeacons {
					ibcMap := ibc.(map[string]interface{})
					var ibcID int
					var ibcName string

					if id, ok := ibcMap["id"].(int); ok {
						ibcID = id
					}
					if name, ok := ibcMap["name"].(string); ok {
						ibcName = name
					}

					ibeacon := jamfpro.MacOSConfigurationProfilesDataSubsetIBeacon{
						IBeacon: jamfpro.MacOSConfigurationProfilesDataSubsetIBeaconItem{
							ID:   ibcID,
							Name: ibcName,
						},
					}
					lim.IBeacons = append(lim.IBeacons, ibeacon)
				}
			}

			scope.Limitations = lim
		}
		// Construct Exclusions
		if exclusions, ok := scopeMap["exclusions"].([]interface{}); ok && len(exclusions) > 0 {
			excMap := exclusions[0].(map[string]interface{})
			var exc jamfpro.MacOSConfigurationProfilesDataSubsetExclusions

			// Safely construct and append Computers in Exclusions
			if computers, ok := excMap["computers"].([]interface{}); ok {
				for _, comp := range computers {
					compMap := comp.(map[string]interface{})
					var compID int
					var compName, compUDID string

					if id, ok := compMap["id"].(int); ok {
						compID = id
					}
					if name, ok := compMap["name"].(string); ok {
						compName = name
					}
					if udid, ok := compMap["udid"].(string); ok {
						compUDID = udid
					}

					computer := jamfpro.MacOSConfigurationProfilesDataSubsetComputer{
						Computer: jamfpro.MacOSConfigurationProfilesDataSubsetComputerItem{
							ID:   compID,
							Name: compName,
							UDID: compUDID,
						},
					}
					exc.Computers = append(exc.Computers, computer)
				}
			}

			// Similar pattern for Buildings, Departments, ComputerGroups, UserGroups, NetworkSegments, IBeacons, JSSUsers, JSSUserGroups
			// Example for Buildings
			if buildings, ok := excMap["buildings"].([]interface{}); ok {
				for _, bld := range buildings {
					bldMap := bld.(map[string]interface{})
					var bldID int
					var bldName string

					if id, ok := bldMap["id"].(int); ok {
						bldID = id
					}
					if name, ok := bldMap["name"].(string); ok {
						bldName = name
					}

					building := jamfpro.MacOSConfigurationProfilesDataSubsetBuilding{
						Building: jamfpro.MacOSConfigurationProfilesDataSubsetBuildingItem{
							ID:   bldID,
							Name: bldName,
						},
					}
					exc.Buildings = append(exc.Buildings, building)
				}
			}

			// Safely construct and append Departments in Exclusions
			if departments, ok := excMap["departments"].([]interface{}); ok {
				for _, dept := range departments {
					deptMap := dept.(map[string]interface{})
					var deptID int
					var deptName string

					if id, ok := deptMap["id"].(int); ok {
						deptID = id
					}
					if name, ok := deptMap["name"].(string); ok {
						deptName = name
					}

					department := jamfpro.MacOSConfigurationProfilesDataSubsetDepartment{
						Department: jamfpro.MacOSConfigurationProfilesDataSubsetDepartmentItem{
							ID:   deptID,
							Name: deptName,
						},
					}
					exc.Departments = append(exc.Departments, department)
				}
			}

			// Safely construct and append ComputerGroups in Exclusions
			if computerGroups, ok := excMap["computer_groups"].([]interface{}); ok {
				for _, grp := range computerGroups {
					grpMap := grp.(map[string]interface{})
					var grpID int
					var grpName string

					if id, ok := grpMap["id"].(int); ok {
						grpID = id
					}
					if name, ok := grpMap["name"].(string); ok {
						grpName = name
					}

					computerGroup := jamfpro.MacOSConfigurationProfilesDataSubsetComputerGroup{
						ComputerGroup: jamfpro.MacOSConfigurationProfilesDataSubsetComputerGroupItem{
							ID:   grpID,
							Name: grpName,
						},
					}
					exc.ComputerGroups = append(exc.ComputerGroups, computerGroup)
				}
			}

			// Similar approach for UserGroups, NetworkSegments, IBeacons, JSSUsers, JSSUserGroups
			// Example for UserGroups
			if userGroups, ok := excMap["user_groups"].([]interface{}); ok {
				for _, grp := range userGroups {
					grpMap := grp.(map[string]interface{})
					var grpID int
					var grpName string

					if id, ok := grpMap["id"].(int); ok {
						grpID = id
					}
					if name, ok := grpMap["name"].(string); ok {
						grpName = name
					}

					userGroup := jamfpro.MacOSConfigurationProfilesDataSubsetUserGroup{
						UserGroup: jamfpro.MacOSConfigurationProfilesDataSubsetUserGroupItem{
							ID:   grpID,
							Name: grpName,
						},
					}
					exc.UserGroups = append(exc.UserGroups, userGroup)
				}
			}

			// Safely construct and append NetworkSegments in Exclusions
			if networkSegments, ok := excMap["network_segments"].([]interface{}); ok {
				for _, seg := range networkSegments {
					segMap := seg.(map[string]interface{})
					var segID int
					var segUID, segName string

					if id, ok := segMap["id"].(int); ok {
						segID = id
					}
					if uid, ok := segMap["uid"].(string); ok {
						segUID = uid
					}
					if name, ok := segMap["name"].(string); ok {
						segName = name
					}

					networkSegment := jamfpro.MacOSConfigurationProfilesDataSubsetNetworkSegment{
						NetworkSegment: jamfpro.MacOSConfigurationProfilesDataSubsetNetworkSegmentItem{
							ID:   segID,
							UID:  segUID,
							Name: segName,
						},
					}
					exc.NetworkSegments = append(exc.NetworkSegments, networkSegment)
				}
			}

			// Safely construct and append IBeacons in Exclusions
			if ibeacons, ok := excMap["ibeacons"].([]interface{}); ok {
				for _, ibc := range ibeacons {
					ibcMap := ibc.(map[string]interface{})
					var ibcID int
					var ibcName string

					if id, ok := ibcMap["id"].(int); ok {
						ibcID = id
					}
					if name, ok := ibcMap["name"].(string); ok {
						ibcName = name
					}

					ibeacon := jamfpro.MacOSConfigurationProfilesDataSubsetIBeacon{
						IBeacon: jamfpro.MacOSConfigurationProfilesDataSubsetIBeaconItem{
							ID:   ibcID,
							Name: ibcName,
						},
					}
					exc.IBeacons = append(exc.IBeacons, ibeacon)
				}
			}

			// Safely construct and append JSSUsers in Exclusions
			if jssUsers, ok := excMap["jss_users"].([]interface{}); ok {
				for _, usr := range jssUsers {
					usrMap := usr.(map[string]interface{})
					var usrID int
					var usrName string

					if id, ok := usrMap["id"].(int); ok {
						usrID = id
					}
					if name, ok := usrMap["name"].(string); ok {
						usrName = name
					}

					jssUser := jamfpro.MacOSConfigurationProfilesDataSubsetJSSUser{
						JSSUser: jamfpro.MacOSConfigurationProfilesDataSubsetJSSUserItem{
							ID:   usrID,
							Name: usrName,
						},
					}
					exc.JSSUsers = append(exc.JSSUsers, jssUser)
				}
			}

			// Safely construct and append JSSUserGroups in Exclusions
			if jssUserGroups, ok := excMap["jss_user_groups"].([]interface{}); ok {
				for _, grp := range jssUserGroups {
					grpMap := grp.(map[string]interface{})
					var grpID int
					var grpName string

					if id, ok := grpMap["id"].(int); ok {
						grpID = id
					}
					if name, ok := grpMap["name"].(string); ok {
						grpName = name
					}

					jssUserGroup := jamfpro.MacOSConfigurationProfilesDataSubsetJSSUserGroup{
						JSSUserGroup: jamfpro.MacOSConfigurationProfilesDataSubsetJSSUserGroupItem{
							ID:   grpID,
							Name: grpName,
						},
					}
					exc.JSSUserGroups = append(exc.JSSUserGroups, jssUserGroup)
				}
			}
			scope.Exclusions = exc
		}
	}
	// Initialize an empty SelfService struct
	selfService := jamfpro.MacOSConfigurationProfilesDataSubsetSelfService{}

	// Retrieve the "self_service" data
	selfServiceData := d.Get("self_service").([]interface{})
	if len(selfServiceData) > 0 {
		ssMap := selfServiceData[0].(map[string]interface{})

		// Safely retrieve and set each field if it exists
		if v, ok := ssMap["install_button_text"].(string); ok {
			selfService.InstallButtonText = v
		}
		if v, ok := ssMap["self_service_description"].(string); ok {
			selfService.SelfServiceDescription = v
		}
		if v, ok := ssMap["force_users_to_view_description"].(bool); ok {
			selfService.ForceUsersToViewDescription = v
		}
		if v, ok := ssMap["feature_on_main_page"].(bool); ok {
			selfService.FeatureOnMainPage = v
		}

		// Safely construct and set SelfServiceIcon
		if iconData, ok := ssMap["self_service_icon"].([]interface{}); ok && len(iconData) > 0 {
			iconMap := iconData[0].(map[string]interface{})
			var iconID int
			var iconURI, iconDataStr string

			if id, ok := iconMap["id"].(int); ok {
				iconID = id
			}
			if uri, ok := iconMap["uri"].(string); ok {
				iconURI = uri
			}
			if data, ok := iconMap["data"].(string); ok {
				iconDataStr = data
			}

			selfService.SelfServiceIcon = jamfpro.MacOSConfigurationProfilesDataSubsetSelfServiceIcon{
				ID:   iconID,
				URI:  iconURI,
				Data: iconDataStr,
			}
		}

		// Safely construct and set SelfServiceCategories
		if categoriesData, ok := ssMap["self_service_categories"].([]interface{}); ok && len(categoriesData) > 0 {
			catMap := categoriesData[0].(map[string]interface{})
			var catID int
			var catName string
			var displayIn, featureIn bool

			if id, ok := catMap["id"].(int); ok {
				catID = id
			}
			if name, ok := catMap["name"].(string); ok {
				catName = name
			}
			if display, ok := catMap["display_in"].(bool); ok {
				displayIn = display
			}
			if feature, ok := catMap["feature_in"].(bool); ok {
				featureIn = feature
			}

			selfService.SelfServiceCategories = jamfpro.MacOSConfigurationProfilesDataSubsetSelfServiceCategories{
				Category: jamfpro.MacOSConfigurationProfilesDataSubsetSelfServiceCategory{
					ID:        catID,
					Name:      catName,
					DisplayIn: displayIn,
					FeatureIn: featureIn,
				},
			}
		}
	}

	// Log the successful construction of the MacOS Config Profile
	log.Printf("[INFO] Successfully constructed macOS Configuration Profile with display name: %s", general.Name)

	return &jamfpro.ResponseMacOSConfigurationProfiles{
		General:     general,
		Scope:       scope,
		SelfService: selfService,
	}

}

// Helper function to generate diagnostics based on the error type
func generateTFDiagsFromHTTPError(err error, d *schema.ResourceData, action string) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceName, exists := d.GetOk("name")
	if !exists {
		resourceName = "unknown"
	}

	// Handle the APIError in the diagnostic
	if apiErr, ok := err.(*http_client.APIError); ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failed to %s the resource with name: %s", action, resourceName),
			Detail:   fmt.Sprintf("API Error (Code: %d): %s", apiErr.StatusCode, apiErr.Message),
		})
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failed to %s the resource with name: %s", action, resourceName),
			Detail:   err.Error(),
		})
	}
	return diags
}

// ResourceJamfProMacOSConfigurationProfilesCreate is responsible for creating a new Jamf Pro macOS Configuration Profile in the remote system.
// The function:
// 1. Constructs the macOS Configuration Profile data using the provided Terraform configuration.
// 2. Calls the API to create the profile in Jamf Pro.
// 3. Updates the Terraform state with the ID of the newly created profile.
// 4. Initiates a read operation to synchronize the Terraform state with the actual state in Jamf Pro.
func ResourceJamfProMacOSConfigurationProfilesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*client.APIClient).Conn
	var diags diag.Diagnostics

	err := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		// Construct the macOS Configuration Profile
		profile := constructJamfProMacOSConfigurationProfile(d)
		if profile == nil {
			return retry.NonRetryableError(fmt.Errorf("failed to construct the macOS Configuration Profile"))
		}

		// Log the details of the profile that is about to be created
		log.Printf("[INFO] Attempting to create macOS Configuration Profile with name: %s", profile.General.Name)

		// Call the API to create the profile and get its ID
		profileID, err := conn.CreateMacOSConfigurationProfile(profile)
		if err != nil {
			log.Printf("[ERROR] Error creating macOS Configuration Profile with name: %s. Error: %s", profile.General.Name, err)
			if apiErr, ok := err.(*http_client.APIError); ok {
				return retry.NonRetryableError(fmt.Errorf("API Error (Code: %d): %s", apiErr.StatusCode, apiErr.Message))
			}
			return retry.RetryableError(err)
		}

		// Log the successfully resource creation
		log.Printf("[INFO] Successfully created macOS Configuration Profile with ID: %d", profileID)

		// Set the ID in the Terraform state
		d.SetId(strconv.Itoa(profileID))

		return nil
	})

	if err != nil {
		return generateTFDiagsFromHTTPError(err, d, "create")
	}

	// Log the ID that was set in the Terraform state
	log.Printf("[INFO] Terraform state was successfully updated with new macOS Configuration Profile with ID: %s", d.Id())

	// Perform a read operation to update the Terraform state
	err = retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		readDiags := ResourceJamfProMacOSConfigurationProfilesRead(ctx, d, meta)
		if len(readDiags) > 0 {
			return retry.RetryableError(fmt.Errorf("failed to read the created resource"))
		}
		return nil
	})

	if err != nil {
		return generateTFDiagsFromHTTPError(err, d, "update state for")
	}

	return diags
}

// ResourceJamfProMacOSConfigurationProfilesRead is responsible for reading the current state of a Jamf Pro macOS Configuration Profile from the remote system.
// The function:
// 1. Fetches the profile's current state using its ID. If it fails then obtain profile's current state using its Name.
// 2. Updates the Terraform state with the fetched data to ensure it accurately reflects the current state in Jamf Pro.
// 3. Handles any discrepancies, such as the profile being deleted outside of Terraform, to keep the Terraform state synchronized.
func ResourceJamfProMacOSConfigurationProfilesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*client.APIClient).Conn
	var diags diag.Diagnostics

	var profile *jamfpro.ResponseMacOSConfigurationProfiles

	// Use the retry function for the read operation
	err := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		// Convert the ID from the Terraform state into an integer to be used for the API request
		profileID, convertErr := strconv.Atoi(d.Id())
		if convertErr != nil {
			return retry.NonRetryableError(fmt.Errorf("failed to parse profile ID: %v", convertErr))
		}

		// Try fetching the profile using the ID
		var apiErr error
		profile, apiErr = conn.GetMacOSConfigurationProfileByID(profileID)
		if apiErr != nil {
			// Handle the APIError
			if apiError, ok := apiErr.(*http_client.APIError); ok {
				return retry.NonRetryableError(fmt.Errorf("API Error (Code: %d): %s", apiError.StatusCode, apiError.Message))
			}
			// If fetching by ID fails, try fetching by Name
			profileName := d.Get("name").(string)
			profile, apiErr = conn.GetMacOSConfigurationProfileByName(profileName)
			if apiErr != nil {
				// Handle the APIError
				if apiError, ok := apiErr.(*http_client.APIError); ok {
					return retry.NonRetryableError(fmt.Errorf("API Error (Code: %d): %s", apiError.StatusCode, apiError.Message))
				}
				return retry.RetryableError(apiErr)
			}
		}
		return nil
	})

	// Handle error from the retry function
	if err != nil {
		// If there's an error while reading the resource, generate diagnostics using the helper function.
		return generateTFDiagsFromHTTPError(err, d, "read")
	}

	// Safely set attributes in the Terraform state
	if err := d.Set("name", profile.General.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("description", profile.General.Description); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("distribution_method", profile.General.DistributionMethod); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("user_removable", profile.General.UserRemovable); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("level", profile.General.Level); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("uuid", profile.General.UUID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("redeploy_on_update", profile.General.RedeployOnUpdate); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// Format XML Payload before setting it in the Terraform state
	formattedPayload, err := formatmacOSConfigurationProfileXMLPayload(profile.General.Payloads)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error formatting XML payload: %s", err))
	}

	if err := d.Set("payloads", formattedPayload); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	siteIDStr := strconv.Itoa(profile.General.Site.ID) // Convert site integer ID to string

	// Only set the site attribute in state if it's not the default value
	if siteIDStr != "-1" || profile.General.Site.Name != "None" {
		if err := d.Set("site", []interface{}{map[string]interface{}{
			"id":   siteIDStr,
			"name": profile.General.Site.Name,
		}}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	} else {
		// If the site is the default value, set it as an empty list to avoid nullifying in the plan
		if err := d.Set("site", []interface{}{}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	categoryIDStr := strconv.Itoa(profile.General.Category.ID) // Convert category integer ID to string

	// Only set the category attribute in state if it's not the default value
	if categoryIDStr != "-1" || profile.General.Category.Name != "No category assigned" {
		if err := d.Set("category", []interface{}{map[string]interface{}{
			"id":   categoryIDStr,
			"name": profile.General.Category.Name,
		}}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	} else {
		// If the category is the default value, set it as an empty list to avoid nullifying in the plan
		if err := d.Set("category", []interface{}{}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	// Check and set each field within the scope attribute
	scopeAttr := map[string]interface{}{}

	scopeAttr["all_computers"] = profile.Scope.AllComputers
	scopeAttr["all_jss_users"] = profile.Scope.AllJSSUsers

	// Safely construct and set Computers within scopeAttr if value not empty.
	var computersList []interface{}
	for _, compItem := range profile.Scope.Computers {
		if compItem.Computer.ID != 0 {
			computerMap := map[string]interface{}{
				"id":   compItem.Computer.ID,
				"name": compItem.Computer.Name,
				"udid": compItem.Computer.UDID,
			}
			// Wrap the computerMap in another map with key 'computer'
			computersList = append(computersList, map[string]interface{}{"computer": []interface{}{computerMap}})
		}
	}
	scopeAttr["computers"] = computersList

	// Safely construct and set Computer Groups
	var computerGroupsList []interface{}
	for _, cgItem := range profile.Scope.ComputerGroups {
		if cgItem.ComputerGroup.ID != 0 {
			computerGroupMap := map[string]interface{}{
				"id":   cgItem.ComputerGroup.ID,
				"name": cgItem.ComputerGroup.Name,
			}
			computerGroupsList = append(computerGroupsList, map[string]interface{}{"computer_group": []interface{}{computerGroupMap}})
		}
	}
	scopeAttr["computer_groups"] = computerGroupsList

	// Safely construct and set jss_users within scope
	var jssUsersList []interface{}
	for _, jssUserItem := range profile.Scope.JSSUsers {
		if jssUserItem.JSSUser.ID != 0 { // Check if ID is not 0
			jssUserMap := map[string]interface{}{
				"id":   jssUserItem.JSSUser.ID,
				"name": jssUserItem.JSSUser.Name,
			}
			jssUsersList = append(jssUsersList, map[string]interface{}{"jss_user": []interface{}{jssUserMap}})
		}
	}
	scopeAttr["jss_users"] = jssUsersList

	// Safely construct and set jss_user_groups within scope
	var jssUserGroupsList []interface{}
	for _, jssUGItem := range profile.Scope.JSSUserGroups {
		if jssUGItem.JSSUserGroup.ID != 0 {
			jssUserGroupMap := map[string]interface{}{
				"id":   jssUGItem.JSSUserGroup.ID,
				"name": jssUGItem.JSSUserGroup.Name,
			}
			// Wrap the jssUserGroupMap in another map with key 'jss_user_group'
			jssUserGroupsList = append(jssUserGroupsList, jssUserGroupMap)
		}
	}
	scopeAttr["jss_user_groups"] = jssUserGroupsList

	// Safely construct and set Buildings within scopeAttr
	var buildingsList []interface{}
	for _, bldgItem := range profile.Scope.Buildings {
		// Only add the building to the list if its ID is not the default (0 in this case)
		if bldgItem.Building.ID != 0 {
			buildingMap := map[string]interface{}{
				"id":   bldgItem.Building.ID,
				"name": bldgItem.Building.Name,
			}
			// Wrap the buildingMap in another map with key 'building'
			buildingsList = append(buildingsList, map[string]interface{}{"building": []interface{}{buildingMap}})
		}
	}
	scopeAttr["buildings"] = buildingsList

	// Safely construct and set Departments if values are not null
	var departmentsList []interface{}
	for _, deptItem := range profile.Scope.Departments {
		if deptItem.Department.ID != 0 {
			departmentMap := map[string]interface{}{
				"id":   deptItem.Department.ID,
				"name": deptItem.Department.Name,
			}
			departmentsList = append(departmentsList, map[string]interface{}{"department": []interface{}{departmentMap}})
		}
	}
	scopeAttr["departments"] = departmentsList

	// Handling Limitations
	limitationsAttr := make(map[string]interface{})

	// Safely construct and set network_segments limitations
	var networkSegmentsList []interface{}
	for _, networkSegmentItem := range profile.Scope.Limitations.NetworkSegments {
		if networkSegmentItem.NetworkSegment.ID != 0 { // Check if ID is not 0
			networkSegmentMap := map[string]interface{}{
				"id":   networkSegmentItem.NetworkSegment.ID,
				"name": networkSegmentItem.NetworkSegment.Name,
				"uid":  networkSegmentItem.NetworkSegment.UID,
			}
			// Wrap the networkSegmentMap in another map with key 'network_segment'
			networkSegmentsList = append(networkSegmentsList, map[string]interface{}{"network_segment": []interface{}{networkSegmentMap}})
		}
	}
	limitationsAttr["network_segments"] = networkSegmentsList

	// Safely construct and set limitations users list
	var usersList []interface{}
	for _, userItem := range profile.Scope.Limitations.Users {
		if userItem.User.ID != 0 {
			userMap := map[string]interface{}{
				"id":   userItem.User.ID,
				"name": userItem.User.Name,
			}
			// Wrap the userMap in another map with key 'user'
			usersList = append(usersList, map[string]interface{}{"user": []interface{}{userMap}})
		}
	}
	limitationsAttr["users"] = usersList

	// Safely construct and set limitations user groups list
	var userGroupsList []interface{}
	for _, userGroupItem := range profile.Scope.Limitations.UserGroups {
		if userGroupItem.UserGroup.ID != 0 {
			userGroupMap := map[string]interface{}{
				"id":   userGroupItem.UserGroup.ID,
				"name": userGroupItem.UserGroup.Name,
			}
			// Wrap the userGroupMap in another map with key 'user_group'
			userGroupsList = append(userGroupsList, map[string]interface{}{"user_group": []interface{}{userGroupMap}})
		}
	}
	limitationsAttr["user_groups"] = userGroupsList

	// Safely construct and set limitations ibeacons list
	var ibeaconsList []interface{}
	for _, ibeaconItem := range profile.Scope.Limitations.IBeacons {
		if ibeaconItem.IBeacon.ID != 0 {
			ibeaconMap := map[string]interface{}{
				"id":   ibeaconItem.IBeacon.ID,
				"name": ibeaconItem.IBeacon.Name,
			}
			// Wrap the ibeaconMap in another map with key 'ibeacon'
			ibeaconsList = append(ibeaconsList, map[string]interface{}{"ibeacon": []interface{}{ibeaconMap}})
		}
	}
	limitationsAttr["ibeacons"] = ibeaconsList

	// After constructing limitationsAttr
	if len(limitationsAttr) > 0 {
		scopeAttr["limitations"] = []interface{}{limitationsAttr}
	} else {
		// Set an empty list to avoid diff if there are no limitations
		scopeAttr["limitations"] = []interface{}{}
	}

	// Handling Exclusions
	exclusionsAttr := make(map[string]interface{})

	// Safely construct and set computers list in exclusions
	var excludedComputersList []interface{}
	for _, computerItem := range profile.Scope.Exclusions.Computers {
		if computerItem.Computer.ID != 0 {
			computerMap := map[string]interface{}{
				"id":   computerItem.Computer.ID,
				"name": computerItem.Computer.Name,
				"udid": computerItem.Computer.UDID,
			}
			// Wrap the computerMap in another map with key 'computer'
			excludedComputersList = append(excludedComputersList, map[string]interface{}{"computer": []interface{}{computerMap}})
		}
	}
	exclusionsAttr["computers"] = excludedComputersList

	// Safely construct and set exclusions computer groups list
	var excludedComputerGroupsList []interface{}
	for _, cgItem := range profile.Scope.Exclusions.ComputerGroups {
		if cgItem.ComputerGroup.ID != 0 {
			computerGroupMap := map[string]interface{}{
				"id":   cgItem.ComputerGroup.ID,
				"name": cgItem.ComputerGroup.Name,
			}
			excludedComputerGroupsList = append(excludedComputerGroupsList, map[string]interface{}{"computer_group": []interface{}{computerGroupMap}})
		}
	}
	exclusionsAttr["computer_groups"] = excludedComputerGroupsList

	// Safely construct and set exclusions jss_users list
	var excludedJSSUsersList []interface{}
	for _, jssUserItem := range profile.Scope.Exclusions.JSSUsers {
		if jssUserItem.JSSUser.ID != 0 { // Check if ID is not 0
			jssUserMap := map[string]interface{}{
				"id":   jssUserItem.JSSUser.ID,
				"name": jssUserItem.JSSUser.Name,
			}
			// Wrap the jssUserMap in another map with key 'jss_user'
			excludedJSSUsersList = append(excludedJSSUsersList, map[string]interface{}{"jss_user": []interface{}{jssUserMap}})
		}
	}
	exclusionsAttr["jss_users"] = excludedJSSUsersList

	// Safely construct and set exclusions jss_user_groups list
	var excludedJSSUserGroupsList []interface{}
	for _, jssUGItem := range profile.Scope.Exclusions.JSSUserGroups {
		if jssUGItem.JSSUserGroup.ID != 0 {
			jssUserGroupMap := map[string]interface{}{
				"id":   jssUGItem.JSSUserGroup.ID,
				"name": jssUGItem.JSSUserGroup.Name,
			}
			// Wrap the jssUserGroupMap in another map with key 'jss_user_group'
			excludedJSSUserGroupsList = append(excludedJSSUserGroupsList, map[string]interface{}{"jss_user_group": []interface{}{jssUserGroupMap}})
		}
	}
	exclusionsAttr["jss_user_groups"] = excludedJSSUserGroupsList

	// Safely construct and set exclusions buildings list
	var excludedBuildingsList []interface{}
	for _, bldgItem := range profile.Scope.Exclusions.Buildings {
		// Only add the building to the list if its ID is not the default (0 in this case)
		if bldgItem.Building.ID != 0 {
			buildingMap := map[string]interface{}{
				"id":   bldgItem.Building.ID,
				"name": bldgItem.Building.Name,
			}
			excludedBuildingsList = append(excludedBuildingsList, map[string]interface{}{"building": []interface{}{buildingMap}})
		}
	}
	exclusionsAttr["buildings"] = excludedBuildingsList

	// Safely construct and set exclusions departments list
	var excludedDepartmentsList []interface{}
	for _, deptItem := range profile.Scope.Exclusions.Departments {
		if deptItem.Department.ID != 0 {
			departmentMap := map[string]interface{}{
				"id":   deptItem.Department.ID,
				"name": deptItem.Department.Name,
			}
			// Wrap the departmentMap in another map with key 'department'
			excludedDepartmentsList = append(excludedDepartmentsList, map[string]interface{}{"department": []interface{}{departmentMap}})
		}
	}
	exclusionsAttr["departments"] = excludedDepartmentsList

	// Safely construct and set exclusions network_segments list
	var excludedNetworkSegmentsList []interface{}
	for _, netSegItem := range profile.Scope.Exclusions.NetworkSegments {
		if netSegItem.NetworkSegment.ID != 0 {
			networkSegmentMap := map[string]interface{}{
				"id":   netSegItem.NetworkSegment.ID,
				"name": netSegItem.NetworkSegment.Name,
				"uid":  netSegItem.NetworkSegment.UID,
			}
			// Wrap the networkSegmentMap in another map with key 'network_segment'
			excludedNetworkSegmentsList = append(excludedNetworkSegmentsList, map[string]interface{}{"network_segment": []interface{}{networkSegmentMap}})
		}
	}
	exclusionsAttr["network_segments"] = excludedNetworkSegmentsList

	// Safely construct and set exclusions user list
	var excludedUsersList []interface{}
	for _, userItem := range profile.Scope.Exclusions.Users {
		if userItem.User.ID != 0 {
			userMap := map[string]interface{}{
				"id":   userItem.User.ID,
				"name": userItem.User.Name,
			}
			// Wrap the userMap in another map with key 'user'
			excludedUsersList = append(excludedUsersList, map[string]interface{}{"user": []interface{}{userMap}})
		}
	}
	exclusionsAttr["users"] = excludedUsersList

	// Safely construct and set exclusions user groups list
	var excludedUserGroupsList []interface{}
	for _, userGroupItem := range profile.Scope.Exclusions.UserGroups {
		if userGroupItem.UserGroup.ID != 0 {
			userGroupMap := map[string]interface{}{
				"id":   userGroupItem.UserGroup.ID,
				"name": userGroupItem.UserGroup.Name,
			}
			// Wrap the userGroupMap in another map with key 'user_group'
			excludedUserGroupsList = append(excludedUserGroupsList, map[string]interface{}{"user_group": []interface{}{userGroupMap}})
		}
	}
	exclusionsAttr["user_groups"] = excludedUserGroupsList

	// Safely construct and set ibeacons list in Exclusions
	var excludedIBeaconsList []interface{}
	for _, ibeaconItem := range profile.Scope.Exclusions.IBeacons {
		if ibeaconItem.IBeacon.ID != 0 {
			ibeaconMap := map[string]interface{}{
				"id":   ibeaconItem.IBeacon.ID,
				"name": ibeaconItem.IBeacon.Name,
			}
			// Wrap the ibeaconMap in another map with key 'ibeacon'
			excludedIBeaconsList = append(excludedIBeaconsList, map[string]interface{}{"ibeacon": []interface{}{ibeaconMap}})
		}
	}
	exclusionsAttr["ibeacons"] = excludedIBeaconsList

	// After constructing exclusionsAttr
	if len(exclusionsAttr) > 0 {
		scopeAttr["exclusions"] = []interface{}{exclusionsAttr}
	} else {
		// Set an empty list to avoid diff if there are no exclusions
		scopeAttr["exclusions"] = []interface{}{}
	}

	// Add the 'scope' to Terraform state
	if err := d.Set("scope", []interface{}{scopeAttr}); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// Setting 'self_service' attribute
	selfServiceAttr := map[string]interface{}{
		"install_button_text":             profile.SelfService.InstallButtonText,
		"self_service_description":        profile.SelfService.SelfServiceDescription,
		"force_users_to_view_description": profile.SelfService.ForceUsersToViewDescription,
		"feature_on_main_page":            profile.SelfService.FeatureOnMainPage,
	}

	// Constructing 'self_service_icon'
	if profile.SelfService.SelfServiceIcon.ID != 0 || profile.SelfService.SelfServiceIcon.URI != "" || profile.SelfService.SelfServiceIcon.Data != "" {
		selfServiceIcon := map[string]interface{}{
			"id":   profile.SelfService.SelfServiceIcon.ID,
			"uri":  profile.SelfService.SelfServiceIcon.URI,
			"data": profile.SelfService.SelfServiceIcon.Data,
		}
		selfServiceAttr["self_service_icon"] = []interface{}{selfServiceIcon}
	}

	// Constructing 'self_service_categories'
	if profile.SelfService.SelfServiceCategories.Category.ID != 0 || profile.SelfService.SelfServiceCategories.Category.Name != "" {
		selfServiceCategory := map[string]interface{}{
			"id":         profile.SelfService.SelfServiceCategories.Category.ID,
			"name":       profile.SelfService.SelfServiceCategories.Category.Name,
			"display_in": profile.SelfService.SelfServiceCategories.Category.DisplayIn,
			"feature_in": profile.SelfService.SelfServiceCategories.Category.FeatureIn,
		}
		selfServiceAttr["self_service_categories"] = []interface{}{selfServiceCategory}
	}

	// Set the 'self_service' attribute in the Terraform state
	if err := d.Set("self_service", []interface{}{selfServiceAttr}); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags

}

// ResourceJamfProMacOSConfigurationProfilesUpdate is responsible for updating an existing Jamf Pro macOS Configuration Profile on the remote system.
func ResourceJamfProMacOSConfigurationProfilesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*client.APIClient).Conn
	var diags diag.Diagnostics

	// Use the retry function for the update operation
	var err error
	err = retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
		// Construct the updated macOS configuration profile
		profile := constructJamfProMacOSConfigurationProfile(d)

		// Convert the ID from the Terraform state into an integer to be used for the API request
		profileID, convertErr := strconv.Atoi(d.Id())
		if convertErr != nil {
			return retry.NonRetryableError(fmt.Errorf("failed to parse profile ID: %v", convertErr))
		}

		// Directly call the API to update the profile
		_, apiErr := conn.UpdateMacOSConfigurationProfileByID(profileID, profile)
		if apiErr != nil {
			// Handle the APIError
			if apiError, ok := apiErr.(*http_client.APIError); ok {
				return retry.NonRetryableError(fmt.Errorf("API Error (Code: %d): %s", apiError.StatusCode, apiError.Message))
			}
			// If the update by ID fails, try updating by name
			profileName := d.Get("name").(string)
			_, apiErr = conn.UpdateMacOSConfigurationProfileByName(profileName, profile)
			if apiErr != nil {
				// Handle the APIError
				if apiError, ok := apiErr.(*http_client.APIError); ok {
					return retry.NonRetryableError(fmt.Errorf("API Error (Code: %d): %s", apiError.StatusCode, apiError.Message))
				}
				return retry.RetryableError(apiErr)
			}
		}
		return nil
	})

	// Handle error from the retry function
	if err != nil {
		// If there's an error while updating the resource, generate diagnostics using the helper function.
		return generateTFDiagsFromHTTPError(err, d, "update")
	}

	// Use the retry function for the read operation to update the Terraform state
	err = retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		readDiags := ResourceJamfProMacOSConfigurationProfilesRead(ctx, d, meta)
		if len(readDiags) > 0 {
			return retry.RetryableError(fmt.Errorf("failed to update the Terraform state for the updated resource"))
		}
		return nil
	})

	// Handle error from the retry function
	if err != nil {
		// If there's an error while updating the Terraform state, generate diagnostics using the helper function.
		return generateTFDiagsFromHTTPError(err, d, "update")
	}

	return diags
}

// ResourceJamfProMacOSConfigurationProfilesDelete is responsible for deleting a macOS Configuration Profile.
func ResourceJamfProMacOSConfigurationProfilesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*client.APIClient).Conn
	var diags diag.Diagnostics

	// Use the retry function for the **DELETE** operation
	err := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		// Convert the ID from the Terraform state into an integer to be used for the API request
		macOSConfigurationProfileID, convertErr := strconv.Atoi(d.Id())
		if convertErr != nil {
			return retry.NonRetryableError(fmt.Errorf("failed to parse department ID: %v", convertErr))
		}

		// Directly call the API to **DELETE** the resource
		apiErr := conn.DeleteMacOSConfigurationProfileByID(macOSConfigurationProfileID)
		if apiErr != nil {
			// If the **DELETE** by ID fails, try deleting by name
			siteName := d.Get("name").(string)
			apiErr = conn.DeleteDepartmentByName(siteName)
			if apiErr != nil {
				return retry.RetryableError(apiErr)
			}
		}
		return nil
	})

	// Handle error from the retry function
	if err != nil {
		// If there's an error while deleting the resource, generate diagnostics using the helper function.
		return generateTFDiagsFromHTTPError(err, d, "delete")
	}

	// Clear the ID from the Terraform state as the resource has been deleted
	d.SetId("")

	return diags
}
