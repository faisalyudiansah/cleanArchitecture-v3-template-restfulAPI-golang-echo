INSERT INTO m_uoms (uom_id, code,description) VALUES
 ('1','PCS','PIECES'),
 ('2','CRT','KARTON');

INSERT INTO m_term_of_payments (term_of_payment_id, days,description) VALUES
 ('1', 30,'Days'),
 ('2', 45,'Days'),
 ('3', 60,'Days');

INSERT INTO m_items (item_id, code,description,uom_id,selling_uom_id,barcode,item_group_id,item_category_id,item_subcategory_id) VALUES
 ('1', 'SAMB-PI-X1','Piattos 68G Barbeque',1,2,'8996196000905',NULL,NULL,NULL),
 ('2', 'SAMB-PI-X2','Piattos 68 G Rumput Laut',1,2,'8996196002954',NULL,NULL,NULL),
 ('3', 'SAMB-PI-X3','Piattos 68 G Sambal Geprek',1,2,'8996196004705',NULL,NULL,NULL),
 ('4', 'SAMB-PI-X4','Piattos 68G Sambal Matah',1,2,'8996196004651',NULL,NULL,NULL),
 ('5', 'SAMB-PI-X5','Piattos 68G Sapi Panggang',1,2,'8996196000950',NULL,NULL,NULL),
 ('6', 'SAMB-PI-X6','Piattos 68G Mystery Flavor',1,2,'8996196005801',NULL,NULL,NULL),
 ('7', 'SAMB-PI-X7','Piattos 120G Sapi Panggang',1,2,'8996196005405',NULL,NULL,NULL),
 ('8', 'SAMB-PI-X8','Piattos 120G Rumput Laut',1,2,'8996196005409',NULL,NULL,NULL);

INSERT INTO m_item_uoms (item_uom_id,item_id,uom_id,uom_value) VALUES
 ('1','1','1',1.000000000),
 ('2','1','2',30.000000000),
 ('3','2','1',1.000000000),
 ('4','2','2',30.000000000),
 ('5','3','1',1.000000000),
 ('6','3','2',30.000000000),
 ('7','4','1',1.000000000),
 ('8','4','2',30.000000000),
 ('9','5','1',1.000000000),
 ('10','5','2',30.000000000);

INSERT INTO m_item_uoms (item_uom_id,item_id,uom_id,uom_value) VALUES
 ('11','6','1',1.000000000),
 ('12','6','2',30.000000000),
 ('13','7','1',1.000000000),
 ('14','7','2',30.000000000),
 ('15','8','1',1.000000000),
 ('16','8','2',16.000000000);

INSERT INTO m_customer_groups (customer_group_id, customer_group_name,credit_limit,credit_amount,term_of_payment_id) VALUES
('5245','PT SUMBER ALFARIA TRIJAYA Tbk',100000000.000000000,100000000.000000000,'1'),
('5663','HARI HARI SWALAYAN',300000000.000000000,300000000.000000000,'1');

INSERT INTO m_customers (customer_id, customer_name,customer_pic,customer_group_id,address,city,postal_code,phone_1,phone_2,npwp,term_of_payment_id) VALUES
 ('3130', 'JAYA PERTAMA ANDALAN PT','Budi','5245','Jl. Raya Pemda Kerad','Bogor','16151','021-6519252',NULL,'31.555.333.6-777.000','1'),
 ('3131', 'SARANA ABADI MAKMUR BERSAMA PT','Andi','5663','Jl. Agung Timur VIII BLOK 0-3/20-21 Sunter Jaya','Tangerang','14350','021-6519259',NULL,'01.330.444.9-038.000','1');

INSERT INTO m_customer_items (customer_item_id,customer_id,item_id,code) VALUES
('1','3130','1','81148813'),
('2','3130','2','80398974'),
('3','3130','3','81139637'),
('4','3130','4','81139620'),
('5','3130','5','81148820'),
('6','3130','6','81211289'),
('7','3130','8','81178520'),
('8','3131','1','81148813'),
('9','3131','2','80398974'),
('10','3131','3','81139637'),
('11','3131','4','81139620'),
('12','3131','5','81148820'),
('13','3131','6','81211289'),
('14','3131','8','81178520');
	
INSERT INTO m_customer_addresses (customer_address_id,customer_id,address_name,address_pic,address,city,postal_code,phone_1,phone_2,is_billing,is_shipto,is_default_billing,is_default_shipto) VALUES
 ('1','3131','SARANA ABADI MAKMUR BERSAMA PT','Andi','Jl. Agung Timur VIII BLOK 0-3/20-21 Sunter Jaya','Jakarta','14350','021-6519259',NULL,1,0,1,0),
 ('2','3131','PT SINARSAHABAT INTIMAKMUR','Andi','Jl. Industri Raya IV BLOK AG No.1','Tangerang','15710','0819-12948118',NULL,0,1,0,1);
 
INSERT INTO `m_areas` (`area_id`, `area_name`, `postal_code`, `area_level`, `area_parent_id`) VALUES
('1', 'Bekasi', '17111', 1, NULL),
('2', 'Tangerang', '15111', 1, NULL),
('3', 'Bogor', '16111', 1, NULL),
('4', 'Jakarta Barat', '11411', 1, NULL),
('5', 'Jakarta Timur', '13411', 1, NULL);

INSERT INTO `m_warehouses` (`warehouse_id`, `warehouse_name`, `city`, `is_bulky`, `fg_warehouse_id`, `fg_warehouse_name`) VALUES
('01', 'Pulogadung', 'Jakarta Timur', 1, '01FP', 'Pulogadung FG Principal'),
('02', 'Pulogebang', 'Jakarta Timur', 1, '02FP', 'Pulogebang FG Principal'),
('03', 'Bekasi', 'Bekasi', 0, '03FP', 'Bekasi FG Principal'),
('04', 'Tangerang', 'Tangerang', 1, '04FP', 'Tangerang FG Principal');

INSERT INTO `m_warehouse_service_areas` (`warehouse_service_area_id`, `warehouse_id`, `area_id`, `bulky_warehouse_id`) VALUES
('1', '3', '1', '02'),
('2', '4', '2', '04'),
('3', '3', '3', '02'),
('4', '1', '4', '04'),
('5', '1', '5', '02');

INSERT INTO `t_inventories` (`inventory_id`, `posting_date`, `posting_time`, `item_id`, `warehouse_id`, `base_qty`, `base_uom_id`, `batch_id`, `stock_type_id`) VALUES
('1', '2024-12-01', '08:00:00', '1', '03', 500, '1', '1', '1'), 
('2', '2024-12-02', '09:30:00', '2', '02', 300, '1', '2', '2'), 
('3', '2024-12-03', '10:45:00', '3', '04', 450, '1', '3', '1'), 
('4', '2024-12-04', '11:15:00', '4', '01', 600, '1', '3', '3'), 
('5', '2024-12-05', '12:00:00', '5', '01', 700, '1', '2', '2'); 

INSERT INTO `m_batches` (`batch_id`, `expiry_date`) VALUES
('1', '2025-12-31'),
('2', '2025-06-30'),
('3', '2026-03-15');

INSERT INTO `m_stock_types` (`stock_type_id`, `stock_type_description`) VALUES
('1', NULL),
('2', NULL),
('3', NULL);

INSERT INTO `r_supplier_items` (`supplier_id`, `item_id`, `client_id`) VALUES
(NULL, '1', 'API'),
(NULL, '2', 'API'),
(NULL, '3', 'API'),
(NULL, '4', 'API'),
(NULL, '5', 'API'),
(NULL, '6', 'API'),
(NULL, '7', 'API'),
(NULL, '8', 'API');