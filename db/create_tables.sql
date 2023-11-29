-- 建立商品資料表
CREATE TABLE `shop_item` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) DEFAULT '' COMMENT '商品名稱',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否刪除 0 為未刪除、1 為已刪除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品管理';

-- 建立競拍資料表
CREATE TABLE `shop_auction` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `item_id` INT(10) NOT NULL COMMENT '商品 ID',
    `init_bid_price` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '起標價',
    `latest_bid_price` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最新標價',
    `hammer_price` INT(10) UNSIGNED NOT NULL DEFAULT '-1' COMMENT '得標價',
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否刪除 0 為未刪除、1 為已刪除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='競拍管理';

-- 建立歷史標價資料表
CREATE TABLE `shop_bidhistory` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `auction_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '競拍 ID',
	`bid_price` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '歷史下標金額',
    
    `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '建立時間',
    `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改時間',
    `deleted_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '刪除時間',
    `is_del` TINYINT(3) UNSIGNED DEFAULT '0' COMMENT '是否刪除 0 為未刪除、1 為已刪除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='競拍歷史記錄';
