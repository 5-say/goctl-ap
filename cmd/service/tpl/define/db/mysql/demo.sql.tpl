DROP TABLE IF EXISTS `{{ .tableName }}`;
CREATE TABLE `{{ .tableName }}` (

  `id`             int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at`     datetime         NOT NULL                COMMENT '创建时间',
  `updated_at`     datetime         DEFAULT NULL            COMMENT '更新时间',
  `deleted_at`     datetime         DEFAULT NULL            COMMENT '删除时间',

  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='';