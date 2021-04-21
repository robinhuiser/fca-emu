// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString},
		{Name: "number", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_opened", Type: field.TypeTime},
		{Name: "date_last_updated", Type: field.TypeTime},
		{Name: "date_closed", Type: field.TypeTime, Nullable: true},
		{Name: "currency_code", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "source", Type: field.TypeString},
		{Name: "interest_reporting", Type: field.TypeBool},
		{Name: "current_balance", Type: field.TypeFloat32},
		{Name: "available_balance", Type: field.TypeFloat32},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "account_branch", Type: field.TypeInt, Nullable: true},
		{Name: "account_product", Type: field.TypeInt, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_branches_branch",
				Columns:    []*schema.Column{AccountsColumns[17]},
				RefColumns: []*schema.Column{BranchesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "accounts_products_product",
				Columns:    []*schema.Column{AccountsColumns[18]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BanksColumns holds the columns for the "banks" table.
	BanksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bank_code", Type: field.TypeString},
		{Name: "bank_name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "swift", Type: field.TypeString},
	}
	// BanksTable holds the schema information for the "banks" table.
	BanksTable = &schema.Table{
		Name:        "banks",
		Columns:     BanksColumns,
		PrimaryKey:  []*schema.Column{BanksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BinaryItemsColumns holds the columns for the "binary_items" table.
	BinaryItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "format", Type: field.TypeString},
		{Name: "length", Type: field.TypeInt},
		{Name: "content", Type: field.TypeBytes},
		{Name: "url", Type: field.TypeString},
		{Name: "transaction_images", Type: field.TypeUUID, Nullable: true},
	}
	// BinaryItemsTable holds the schema information for the "binary_items" table.
	BinaryItemsTable = &schema.Table{
		Name:       "binary_items",
		Columns:    BinaryItemsColumns,
		PrimaryKey: []*schema.Column{BinaryItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "binary_items_transactions_images",
				Columns:    []*schema.Column{BinaryItemsColumns[5]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BranchesColumns holds the columns for the "branches" table.
	BranchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "branch_code", Type: field.TypeString},
		{Name: "street_number", Type: field.TypeString},
		{Name: "street_name", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "state", Type: field.TypeString, Size: 2},
		{Name: "zip", Type: field.TypeString},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "bank_branches", Type: field.TypeInt, Nullable: true},
	}
	// BranchesTable holds the schema information for the "branches" table.
	BranchesTable = &schema.Table{
		Name:       "branches",
		Columns:    BranchesColumns,
		PrimaryKey: []*schema.Column{BranchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "branches_banks_branches",
				Columns:    []*schema.Column{BranchesColumns[9]},
				RefColumns: []*schema.Column{BanksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"CREDIT", "DEBIT"}},
		{Name: "number", Type: field.TypeString},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "expiry_date", Type: field.TypeTime},
		{Name: "holder_name", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"LOCKED", "BLOCKED", "OK"}},
		{Name: "url", Type: field.TypeString},
		{Name: "card_network", Type: field.TypeInt, Nullable: true},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_card_networks_network",
				Columns:    []*schema.Column{CardsColumns[8]},
				RefColumns: []*schema.Column{CardNetworksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CardNetworksColumns holds the columns for the "card_networks" table.
	CardNetworksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
	}
	// CardNetworksTable holds the schema information for the "card_networks" table.
	CardNetworksTable = &schema.Table{
		Name:        "card_networks",
		Columns:     CardNetworksColumns,
		PrimaryKey:  []*schema.Column{CardNetworksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EntitiesColumns holds the columns for the "entities" table.
	EntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "firstname", Type: field.TypeString, Nullable: true},
		{Name: "lastname", Type: field.TypeString, Nullable: true},
		{Name: "fullname", Type: field.TypeString, Nullable: true},
		{Name: "date_of_birth", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"PERSON", "ORGANIZATION", "CORPORATE"}},
		{Name: "last_login_date", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
	}
	// EntitiesTable holds the schema information for the "entities" table.
	EntitiesTable = &schema.Table{
		Name:        "entities",
		Columns:     EntitiesColumns,
		PrimaryKey:  []*schema.Column{EntitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EntityAddressesColumns holds the columns for the "entity_addresses" table.
	EntityAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "country", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "postal_code", Type: field.TypeString},
		{Name: "state", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"BUSINESS", "PRIVATE", "MAILBOX"}},
		{Name: "line1", Type: field.TypeString},
		{Name: "line2", Type: field.TypeString, Nullable: true},
		{Name: "line3", Type: field.TypeString, Nullable: true},
		{Name: "primary", Type: field.TypeBool},
		{Name: "entity_addresses", Type: field.TypeUUID, Nullable: true},
	}
	// EntityAddressesTable holds the schema information for the "entity_addresses" table.
	EntityAddressesTable = &schema.Table{
		Name:       "entity_addresses",
		Columns:    EntityAddressesColumns,
		PrimaryKey: []*schema.Column{EntityAddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entity_addresses_entities_addresses",
				Columns:    []*schema.Column{EntityAddressesColumns[10]},
				RefColumns: []*schema.Column{EntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EntityContactPointsColumns holds the columns for the "entity_contact_points" table.
	EntityContactPointsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "prefix", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"SMS", "EMAIL", "PHONE", "WHATSAPP", "SKYPE"}},
		{Name: "suffix", Type: field.TypeString, Nullable: true},
		{Name: "value", Type: field.TypeString},
		{Name: "entity_contact_points", Type: field.TypeUUID, Nullable: true},
	}
	// EntityContactPointsTable holds the schema information for the "entity_contact_points" table.
	EntityContactPointsTable = &schema.Table{
		Name:       "entity_contact_points",
		Columns:    EntityContactPointsColumns,
		PrimaryKey: []*schema.Column{EntityContactPointsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entity_contact_points_entities_contactPoints",
				Columns:    []*schema.Column{EntityContactPointsColumns[6]},
				RefColumns: []*schema.Column{EntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EntityTaxInformationsColumns holds the columns for the "entity_tax_informations" table.
	EntityTaxInformationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"SSN"}},
		{Name: "tax_id", Type: field.TypeString},
		{Name: "entity_tax_specifications", Type: field.TypeUUID, Nullable: true},
	}
	// EntityTaxInformationsTable holds the schema information for the "entity_tax_informations" table.
	EntityTaxInformationsTable = &schema.Table{
		Name:       "entity_tax_informations",
		Columns:    EntityTaxInformationsColumns,
		PrimaryKey: []*schema.Column{EntityTaxInformationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entity_tax_informations_entities_taxSpecifications",
				Columns:    []*schema.Column{EntityTaxInformationsColumns[3]},
				RefColumns: []*schema.Column{EntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PreferencesColumns holds the columns for the "preferences" table.
	PreferencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "account_preferences", Type: field.TypeUUID, Nullable: true},
		{Name: "entity_preferences", Type: field.TypeUUID, Nullable: true},
	}
	// PreferencesTable holds the schema information for the "preferences" table.
	PreferencesTable = &schema.Table{
		Name:       "preferences",
		Columns:    PreferencesColumns,
		PrimaryKey: []*schema.Column{PreferencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "preferences_accounts_preferences",
				Columns:    []*schema.Column{PreferencesColumns[3]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "preferences_entities_preferences",
				Columns:    []*schema.Column{PreferencesColumns[4]},
				RefColumns: []*schema.Column{EntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"SAVING", "LOAN", "DEPOSIT", "CHECKING", "INVESTMENT"}},
		{Name: "type_name", Type: field.TypeString},
		{Name: "sub_type", Type: field.TypeString},
		{Name: "sub_type_name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:        "products",
		Columns:     ProductsColumns,
		PrimaryKey:  []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RoutingNumbersColumns holds the columns for the "routing_numbers" table.
	RoutingNumbersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WIRE", "ABA"}},
		{Name: "account_routingnumbers", Type: field.TypeUUID, Nullable: true},
	}
	// RoutingNumbersTable holds the schema information for the "routing_numbers" table.
	RoutingNumbersTable = &schema.Table{
		Name:       "routing_numbers",
		Columns:    RoutingNumbersColumns,
		PrimaryKey: []*schema.Column{RoutingNumbersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routing_numbers_accounts_routingnumbers",
				Columns:    []*schema.Column{RoutingNumbersColumns[3]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "sequence_in_day", Type: field.TypeInt, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "POSTED"}},
		{Name: "executed_amount", Type: field.TypeFloat64},
		{Name: "executed_currency_code", Type: field.TypeString, Size: 3},
		{Name: "exchange_rate", Type: field.TypeFloat64},
		{Name: "originating_amount", Type: field.TypeFloat64},
		{Name: "originating_currency_code", Type: field.TypeString, Size: 3},
		{Name: "direction", Type: field.TypeEnum, Enums: []string{"DEBIT", "CREDIT"}},
		{Name: "running_balance", Type: field.TypeFloat64},
		{Name: "created_date", Type: field.TypeTime},
		{Name: "posted_date", Type: field.TypeTime},
		{Name: "executed_date", Type: field.TypeTime},
		{Name: "updated_date", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "memo", Type: field.TypeString, Nullable: true},
		{Name: "group", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "main_category", Type: field.TypeString, Nullable: true},
		{Name: "sub_category", Type: field.TypeString, Nullable: true},
		{Name: "check_number", Type: field.TypeString, Nullable: true},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "merchant_code", Type: field.TypeString, Nullable: true},
		{Name: "reversal", Type: field.TypeBool},
		{Name: "reversal_for", Type: field.TypeString, Nullable: true},
		{Name: "reversed", Type: field.TypeBool},
		{Name: "reversed_by", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "account_transactions", Type: field.TypeUUID, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_accounts_transactions",
				Columns:    []*schema.Column{TransactionsColumns[29]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AccountOwnersColumns holds the columns for the "account_owners" table.
	AccountOwnersColumns = []*schema.Column{
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "entity_id", Type: field.TypeUUID},
	}
	// AccountOwnersTable holds the schema information for the "account_owners" table.
	AccountOwnersTable = &schema.Table{
		Name:       "account_owners",
		Columns:    AccountOwnersColumns,
		PrimaryKey: []*schema.Column{AccountOwnersColumns[0], AccountOwnersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_owners_account_id",
				Columns:    []*schema.Column{AccountOwnersColumns[0]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "account_owners_entity_id",
				Columns:    []*schema.Column{AccountOwnersColumns[1]},
				RefColumns: []*schema.Column{EntitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		BanksTable,
		BinaryItemsTable,
		BranchesTable,
		CardsTable,
		CardNetworksTable,
		EntitiesTable,
		EntityAddressesTable,
		EntityContactPointsTable,
		EntityTaxInformationsTable,
		PreferencesTable,
		ProductsTable,
		RoutingNumbersTable,
		TransactionsTable,
		AccountOwnersTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = BranchesTable
	AccountsTable.ForeignKeys[1].RefTable = ProductsTable
	BinaryItemsTable.ForeignKeys[0].RefTable = TransactionsTable
	BranchesTable.ForeignKeys[0].RefTable = BanksTable
	CardsTable.ForeignKeys[0].RefTable = CardNetworksTable
	EntityAddressesTable.ForeignKeys[0].RefTable = EntitiesTable
	EntityContactPointsTable.ForeignKeys[0].RefTable = EntitiesTable
	EntityTaxInformationsTable.ForeignKeys[0].RefTable = EntitiesTable
	PreferencesTable.ForeignKeys[0].RefTable = AccountsTable
	PreferencesTable.ForeignKeys[1].RefTable = EntitiesTable
	RoutingNumbersTable.ForeignKeys[0].RefTable = AccountsTable
	TransactionsTable.ForeignKeys[0].RefTable = AccountsTable
	AccountOwnersTable.ForeignKeys[0].RefTable = AccountsTable
	AccountOwnersTable.ForeignKeys[1].RefTable = EntitiesTable
}
