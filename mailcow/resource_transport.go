package mailcow

func resourceTransport() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTransportCreate,
		ReadContext:   resourceTransportRead,
		UpdateContext: resourceTransportUpdate,
		DeleteContext: resourceTransportDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceTransportImport,
		},

		Schema: map[string]*schema.Schema{
			// mailcow
			"active": {
				Type:        schema.TypeBool,
				Description: "enable or disable transport",
				Default:     true,
				Optional:    true,
			},
			"destination": {
				Type:        schema.TypeString,
				Description: "destination of the transport",
				Required:    true,
			},
			"nexthop": {
				Type:        schema.TypeString,
				Description: "next hop for messages",
				Required:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "username for the transport",
				Optional:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Description: "password for the transport",
				Optional:    true,
			},
		},
	}
}

func resourceTransportImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}

func resourceTransportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*APIClient)

	mailcowCreateRequest := api.NewCreateTransportRequest()
	createRequestSet(mailcowCreateRequest, resourceTransport(), d, nil, nil)

	err := mailcowCreate(ctx, resourceTransport(), d, "", nil, nil, mailcowCreateRequest, c)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := response.GetTransportId()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*id)

	return diags
}
