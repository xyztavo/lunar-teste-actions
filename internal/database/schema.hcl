table "products" {
  schema = schema.main
  column "sku" {
    null = false
    type = text
  }
  column "title" {
    null = false
    type = text
  }
  column "priceCents" {
    null = false
    type = integer
  }
  primary_key {
    columns = [column.sku]
  }
}
schema "main" {
}
