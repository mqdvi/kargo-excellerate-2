ALTER TABLE `districts` CHANGE COLUMN `id` `id` INT(11) NOT NULL AUTO_INCREMENT FIRST;
ALTER TABLE `drivers` CHANGE COLUMN `id` `id` INT(11) NOT NULL AUTO_INCREMENT FIRST;
ALTER TABLE `shipments` CHANGE COLUMN `id` `id` INT(11) NOT NULL AUTO_INCREMENT FIRST;
ALTER TABLE `truck_types` CHANGE COLUMN `id` `id` INT(11) NOT NULL AUTO_INCREMENT FIRST;
ALTER TABLE `trucks` CHANGE COLUMN `id` `id` INT(11) NOT NULL AUTO_INCREMENT FIRST;

ALTER TABLE `trucks`
	CHANGE COLUMN `trucks_type_id` `truck_type_id` INT(11) NULL DEFAULT NULL AFTER `license_number`,
	DROP INDEX `trucks_type_id`,
	ADD INDEX `trucks_type_id` (`truck_type_id`) USING BTREE;
SELECT `DEFAULT_COLLATION_NAME` FROM `information_schema`.`SCHEMATA` WHERE `SCHEMA_NAME`='db_kargo_2';