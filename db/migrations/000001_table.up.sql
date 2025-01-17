-- masters.m_item_groups definition

CREATE TABLE `m_item_groups` (
  `item_group_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(100) NOT NULL,
  `description` longtext DEFAULT NULL,
  PRIMARY KEY (`item_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `m_term_of_payments` (
  `term_of_payment_id` varchar(100) NOT NULL,
  `days` int(10) unsigned NOT NULL,
  `description` varchar(100) NOT NULL,
  PRIMARY KEY (`term_of_payment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_uoms definition

CREATE TABLE `m_uoms` (
  `uom_id` varchar(100) NOT NULL,
  `code` varchar(100) DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`uom_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_customer_groups definition

CREATE TABLE `m_customer_groups` (
  `customer_group_id` varchar(100) NOT NULL,
  `customer_group_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `credit_limit` decimal(21,9) NOT NULL DEFAULT 0.000000000,
  `credit_amount` decimal(21,9) NOT NULL DEFAULT 0.000000000,
  `term_of_payment_id` varchar(100) NOT NULL,
  PRIMARY KEY (`customer_group_id`),
  KEY `m_customer_groups_m_term_of_payments_FK` (`term_of_payment_id`),
  CONSTRAINT `m_customer_groups_m_term_of_payments_FK` FOREIGN KEY (`term_of_payment_id`) REFERENCES `m_term_of_payments` (`term_of_payment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_customers definition

CREATE TABLE `m_customers` (
  `customer_id` varchar(100) NOT NULL,
  `customer_name` longtext NOT NULL,
  `customer_pic` varchar(100) DEFAULT NULL,
  `customer_group_id` varchar(100) NOT NULL,
  `address` longtext NOT NULL,
  `city` varchar(100) NOT NULL,
  `postal_code` varchar(8) NOT NULL,
  `phone_1` varchar(16) NOT NULL,
  `phone_2` varchar(16) DEFAULT NULL,
  `npwp` varchar(100) DEFAULT NULL,
  `term_of_payment_id` varchar(100) NOT NULL,
  PRIMARY KEY (`customer_id`),
  KEY `m_customers_m_customer_groups_FK` (`customer_group_id`),
  KEY `m_customers_m_term_of_payments_FK` (`term_of_payment_id`),
  CONSTRAINT `m_customers_m_customer_groups_FK` FOREIGN KEY (`customer_group_id`) REFERENCES `m_customer_groups` (`customer_group_id`),
  CONSTRAINT `m_customers_m_term_of_payments_FK` FOREIGN KEY (`term_of_payment_id`) REFERENCES `m_term_of_payments` (`term_of_payment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_item_categories definition

CREATE TABLE `m_item_categories` (
  `item_category_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `item_group_id` bigint(20) unsigned DEFAULT NULL,
  `code` varchar(100) NOT NULL,
  `description` longtext DEFAULT NULL,
  PRIMARY KEY (`item_category_id`),
  KEY `m_item_categories_m_item_groups_FK` (`item_group_id`),
  CONSTRAINT `m_item_categories_m_item_groups_FK` FOREIGN KEY (`item_group_id`) REFERENCES `m_item_groups` (`item_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_item_subcategories definition

CREATE TABLE `m_item_subcategories` (
  `item_subcategory_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `item_group_id` bigint(20) unsigned DEFAULT NULL,
  `item_category_id` bigint(20) unsigned NOT NULL,
  `code` varchar(100) NOT NULL,
  `description` longtext DEFAULT NULL,
  PRIMARY KEY (`item_subcategory_id`),
  KEY `m_item_subcategories_m_item_groups_FK` (`item_group_id`),
  KEY `m_item_subcategories_m_item_categories_FK` (`item_category_id`),
  CONSTRAINT `m_item_subcategories_m_item_categories_FK` FOREIGN KEY (`item_category_id`) REFERENCES `m_item_categories` (`item_category_id`),
  CONSTRAINT `m_item_subcategories_m_item_groups_FK` FOREIGN KEY (`item_group_id`) REFERENCES `m_item_groups` (`item_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_items definition

CREATE TABLE `m_items` (
  `item_id` varchar(100) NOT NULL,
  `code` varchar(16) NOT NULL,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `uom_id` varchar(100) DEFAULT NULL,
  `selling_uom_id` varchar(100) DEFAULT NULL,
  `barcode` varchar(100) NOT NULL,
  `item_group_id` varchar(100) DEFAULT NULL,
  `item_category_id` varchar(100) DEFAULT NULL,
  `item_subcategory_id` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`item_id`),
  KEY `m_items_uom_id_IDX` (`uom_id`) USING HASH,
  KEY `m_items_selling_uom_id_IDX` (`selling_uom_id`) USING HASH,
  KEY `m_items_barcode_IDX` (`barcode`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_uom_conversions definition

CREATE TABLE `m_uom_conversions` (
  `uom_conversion_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `item_id` varchar(100) NOT NULL,
  `conversion` decimal(12,9) DEFAULT NULL,
  `to_uom_id` varchar(100) NOT NULL,
  `from_uom_id` varchar(100) NOT NULL,
  PRIMARY KEY (`uom_conversion_id`),
  KEY `m_uom_conversions_item_id_IDX` (`item_id`) USING HASH,
  KEY `m_uom_conversions_to_uom_id_IDX` (`to_uom_id`,`from_uom_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_customer_addresses definition

CREATE TABLE `m_customer_addresses` (
  `customer_address_id` varchar(100) NOT NULL,
  `customer_id` varchar(100) NOT NULL,
  `address_name` varchar(100) NOT NULL,
  `address_pic` varchar(100) DEFAULT NULL,
  `address` varchar(100) NOT NULL,
  `city` varchar(100) NOT NULL,
  `postal_code` varchar(100) NOT NULL,
  `phone_1` varchar(100) NOT NULL,
  `phone_2` varchar(100) DEFAULT NULL,
  `is_billing` tinyint(4) DEFAULT 1,
  `is_shipto` tinyint(4) DEFAULT 1,
  `is_default_billing` tinyint(4) NOT NULL DEFAULT 0,
  `is_default_shipto` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`customer_address_id`),
  KEY `m_customer_addresses_m_customers_FK` (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_customer_items definition

CREATE TABLE `m_customer_items` (
  `customer_item_id` varchar(100) NOT NULL,
  `customer_id` varchar(100) NOT NULL,
  `item_id` varchar(100) NOT NULL,
  `code` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`customer_item_id`),
  KEY `m_customer_items_customer_id_IDX` (`customer_id`) USING HASH,
  KEY `m_customer_items_code_IDX` (`code`) USING HASH,
  KEY `m_customer_items_customer_id_code_IDX` (`customer_id`,`code`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_item_uoms definition

CREATE TABLE `m_item_uoms` (
  `item_uom_id` varchar(100) NOT NULL,
  `item_id` varchar(100) NOT NULL,
  `uom_id` varchar(100) NOT NULL,
  `uom_value` decimal(12,9) NOT NULL,
  PRIMARY KEY (`item_uom_id`),
  KEY `m_item_uoms_item_id_index` (`item_id`) USING HASH,
  KEY `m_item_uoms_uom_id_IDX` (`uom_id`) USING HASH,
  KEY `m_item_uoms_item_id_uom_id_index` (`item_id`,`uom_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- masters.m_areas definition

CREATE TABLE `m_areas` (
  `area_id` varchar(100) NOT NULL,
  `area_name` longtext NOT NULL,
  `postal_code` varchar(100) DEFAULT NULL,
  `area_level` tinyint(4) NOT NULL DEFAULT 1,
  `area_parent_id` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`area_id`),
  KEY `m_areas_area_parent_id_IDX` (`area_parent_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_warehouse_service_areas definition

CREATE TABLE `m_warehouse_service_areas` (
  `warehouse_service_area_id` varchar(100) NOT NULL,
  `warehouse_id` varchar(100) NOT NULL,
  `area_id` varchar(100) NOT NULL,
  `bulky_warehouse_id` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`warehouse_service_area_id`),
  KEY `m_warehouse_service_areas_warehouse_id_IDX` (`warehouse_id`) USING HASH,
  KEY `m_warehouse_service_areas_area_id_IDX` (`area_id`) USING HASH,
  KEY `m_warehouse_service_areas_bulky_warehouse_id_IDX` (`bulky_warehouse_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- masters.m_warehouses definition

CREATE TABLE `m_warehouses` (
  `warehouse_id` varchar(100) NOT NULL,
  `warehouse_name` longtext NOT NULL,
  `city` longtext DEFAULT NULL,
  `is_bulky` tinyint(4) NOT NULL DEFAULT 0,
  `fg_warehouse_id` varchar(100) NOT NULL COMMENT 'Freegoods Warehouse ID',
  `fg_warehouse_name` longtext NOT NULL,
  PRIMARY KEY (`warehouse_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- masters.t_inventories definition

CREATE TABLE `t_inventories` (
  `inventory_id` varchar(100) NOT NULL,
  `posting_date` date NOT NULL,
  `posting_time` time NOT NULL,
  `item_id` varchar(100) DEFAULT NULL,
  `warehouse_id` varchar(100) DEFAULT NULL,
  `base_qty` decimal(10,0) NOT NULL,
  `base_uom_id` varchar(100) NOT NULL,
  `batch_id` varchar(100) DEFAULT NULL,
  `stock_type_id` varchar(100) NOT NULL,
  PRIMARY KEY (`inventory_id`),
  KEY `t_inventories_posting_date_IDX` (`posting_date`,`posting_time`) USING BTREE,
  KEY `t_inventories_item_id_IDX` (`item_id`) USING HASH,
  KEY `t_inventories_warehouse_id_IDX` (`warehouse_id`) USING HASH,
  KEY `t_inventories_batch_id_IDX` (`batch_id`) USING HASH,
  KEY `t_inventories_stock_type_id_IDX` (`stock_type_id`) USING HASH,
  KEY `t_inventories_item_warehouse_id_IDX` (`item_id`,`warehouse_id`) USING BTREE,
  KEY `t_inventories_item_warehouse_batch_id_IDX` (`item_id`,`warehouse_id`,`stock_type_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_batches definition

CREATE TABLE `m_batches` (
  `batch_id` varchar(100) NOT NULL,
  `expiry_date` date NOT NULL,
  PRIMARY KEY (`batch_id`),
  KEY `m_batches_expiry_date_IDX` (`expiry_date`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- masters.m_stock_types definition

CREATE TABLE `m_stock_types` (
  `stock_type_id` varchar(100) NOT NULL,
  `stock_type_description` longtext DEFAULT NULL,
  PRIMARY KEY (`stock_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- masters.r_supplier_items definition

CREATE TABLE `r_supplier_items` (
  `supplier_id` varchar(100),
  `item_id` varchar(100) NOT NULL,
  `client_id` varchar(100) NOT NULL,
  KEY `r_supplier_items_supplier_id_IDX` (`supplier_id`) USING HASH,
  KEY `r_supplier_items_item_id_IDX` (`item_id`) USING HASH,
  KEY `r_supplier_items_client_id_IDX` (`client_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

