-- -------------------------------------
-- Table structure for ecom_user
-- -------------------------------------
DROP TABLE IF EXISTS `ecom_user`;
CREATE TABLE IF NOT EXISTS `ecom_user` (
   `usr_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
   `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
   `usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone Number',
   `usr_username` varchar(30) NOT NULL DEFAULT '' COMMENT 'Username',
   `usr_password` varchar(64) NOT NULL DEFAULT '' COMMENT 'Password',
   `usr_created_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Creation Time',
   `usr_updated_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Update Time',
   `usr_create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Creation IP',
   `usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last Login IP',
   `usr_last_login_at` int(11) NOT NULL DEFAULT '0' COMMENT 'Last Login Time',
   `usr_login_times` int(11) NOT NULL DEFAULT '0' COMMENT 'Login Times',
   `usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status 1:enable, 0:disable, -1: deleted',
   PRIMARY KEY (`usr_id`),
   KEY `idx_email` (`usr_email`),
   KEY `idx_phone` (`usr_phone`),
   KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account';
-- -------------------------------------
-- Table structure for pre_go_acc_user_verify_9999
-- -------------------------------------
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_verify_9999` (
  `verify_id` int NOT NULL AUTO_INCREMENT,
  `verify_otp` varchar(6) NOT NULL,
  `verify_key` varchar(255) NOT NULL,
  `verify_key_hash` varchar(255) NOT NULL,
  `verify_type` int DEFAULT '1',
  `is_verified` int DEFAULT '0',
  `is_deleted` int DEFAULT '0',
  `verify_created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `verify_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`verify_id`),
  UNIQUE KEY `unique_verify_key` (`verify_key`),
  KEY `idx_verify_otp` (`verify_otp`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='account_user_verify';