/*创建数据库*/
DROP DATABASE IF EXISTS `ll_cms`;
CREATE DATABASE IF NOT EXISTS `ll_cms`;
USE `ll_cms`;

DROP TABLE IF EXISTS `ll_admin`;
CREATE TABLE `ll_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员ID 主键',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '登录用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '登录密码',
  `realname` varchar(32) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(16) NOT NULL DEFAULT '' COMMENT '手机',
  `description` varchar(128) NOT NULL DEFAULT '' COMMENT '描述',
  `login_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '最后一次登陆时间',
  `login_count` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1=正常，0=禁用',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级用户的ID  0=顶级用户',
  `is_super` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是超级管理员,超级管理员登录后不受权限约束',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 1=删除  0=不删除',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理员信息表';

DROP TABLE IF EXISTS `ll_role`;
CREATE TABLE `ll_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID 主键',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '角色名称',
  `sort` int(11) NOT NULL DEFAULT '100' COMMENT '排序',
  `description` varchar(128) NOT NULL DEFAULT '' COMMENT '角色描述',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '启用状态：1=启用；0=关闭',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '保留字段，修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='角色信息表-用于管理员';

DROP TABLE IF EXISTS `ll_admin_role`;
CREATE TABLE `ll_admin_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员角色ID 主键',
  `admin_id` int(11) NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理员角色表';


DROP TABLE IF EXISTS `ll_article`;
CREATE TABLE `ll_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章ID 主键',
  `cat_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章分类 0=未分类',
  `cat_id_link` varchar(64) NOT NULL DEFAULT '' COMMENT '分类ID链，方便根据大类查询大类和大类的所有子类下的文章',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `platform` varchar(16) NOT NULL DEFAULT '' COMMENT '发布人所属的平台名称',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '文章标题',
  `title_style` varchar(128) NOT NULL DEFAULT '' COMMENT '标题样式',
  `keyword` varchar(32) NOT NULL DEFAULT '' COMMENT '文章关键词',
  `is_open` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否开放：1=开放， 0=不开放',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态：0=未审核；1=已审核；2=审核未通过',
  `att_url` varchar(128) NOT NULL DEFAULT '' COMMENT '文章缩略图URL',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '权重，默认为文章ID',
  `description` varchar(1255) NOT NULL DEFAULT '' COMMENT '文章描述',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  PRIMARY KEY (`id`),
  KEY `idx_cat_id` (`cat_id`) USING BTREE, /*创建索引*/
  KEY `idx_sort` (`sort`) USING BTREE,
  KEY `idx_cat_id_link` (`cat_id_link`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COMMENT='文章基本信息表';

DROP TABLE IF EXISTS `ll_article_content`;
CREATE TABLE `ll_article_content` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章ID 主键',
  `content` varchar(20480) NOT NULL COMMENT '文章内容',
  `author` varchar(50) NOT NULL DEFAULT '' COMMENT '文章作者',
  `copyfrom` varchar(400) NOT NULL DEFAULT '' COMMENT '文章来源',
  `vote_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '投票ID',
  `review` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否允许评论：0=不允许；1=允许；',
  `review_num` int(10) unsigned DEFAULT '0' COMMENT '评论次数',
  `view_num` int(11) NOT NULL DEFAULT '0' COMMENT '浏览次数',
  `like_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '赞的次数',
  `unlike_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '踩的次数',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COMMENT='文章详情表';

DROP TABLE IF EXISTS `ll_article_cat`;
CREATE TABLE `ll_article_cat` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID,主键',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '分类名称',
  `keyword` varchar(100) NOT NULL DEFAULT '' COMMENT '分类关键词',
  `description` varchar(400) NOT NULL DEFAULT '' COMMENT '分类描述',
  `sort` int(10) unsigned NOT NULL DEFAULT '100' COMMENT '排序',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级分类ID',
  `comment` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否允许评论: 1=允许，0=不允许',
  `is_open` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否公开：0=隐藏；2=公开；',
  `rank` tinyint(4) NOT NULL DEFAULT '0' COMMENT '分类等级',
  `is_system` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否为系统分类：0=普通分类；1=系统分类（不允许删除）',
  `att_url` varchar(255) NOT NULL DEFAULT '' COMMENT '分类图标',
  `review` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否允许评论：0=允许；1=允许；3=默认根据分类设置为准',
  `id_link` varchar(64) NOT NULL DEFAULT '' COMMENT '分类ID链，英文逗号分隔的ID序列，记录分类的上下级关系',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='文章分类表';

DROP TABLE IF EXISTS `ll_article_review`;
CREATE TABLE `ll_article_review` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID 主键',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级评论的ID，此值大于0表示此贴是对帖子的回复',
  `parent_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级帖主的用户ID：0=没有上级用户',
  `parent_username` varchar(32) NOT NULL DEFAULT '' COMMENT '上级贴主的用户名',
  `parent_user_logo` varchar(255) NOT NULL DEFAULT '' COMMENT '上级贴主的头像',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `username` varchar(32) NOT NULL DEFAULT '',
  `user_logo` varchar(128) NOT NULL,
  `article_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
  `article_title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` varchar(512) NOT NULL DEFAULT '' COMMENT '评论内容',
  `like_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '赞的次数',
  `unlike_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '踩的次数',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态：0=未审核；1=已审核；2=审核未通过',
  `client_ip` varchar(39) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_article_id` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='文章评论表';

DROP TABLE IF EXISTS `ll_email`;
CREATE TABLE `ll_email` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '邮件记录ID 主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '发送状态：0=未发送；1=发送成功；2=发送失败',
  `type` varchar(32) NOT NULL DEFAULT '' COMMENT '发送邮件的业务类型',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT 'email',
  `subject` varchar(32) NOT NULL DEFAULT '' COMMENT '邮件标题',
  `body` varchar(10240) NOT NULL DEFAULT '' COMMENT '邮件详情',
  `return_info` varchar(255) NOT NULL DEFAULT '' COMMENT '提示信息',
  `client_ip` varchar(39) NOT NULL DEFAULT '' COMMENT '客户端IP',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='邮件发送记录表';

DROP TABLE IF EXISTS `ll_link`;
CREATE TABLE `ll_link` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '链接ID 主键',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '品牌名称',
  `sort` int(11) NOT NULL DEFAULT '100' COMMENT '排序',
  `tag` varchar(32) NOT NULL DEFAULT '' COMMENT '链接标签，用于分类查找',
  `att_url` varchar(128) NOT NULL DEFAULT '' COMMENT 'logo地址 只对图片链接有效',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT '品牌网址',
  `remark` varchar(512) NOT NULL COMMENT '品牌描述',
  `is_open` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否公开',
  `start_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '开始日期，1970-01-01 08:00:01表示不限',
  `end_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '截止日期，1970-01-01 08:00:01表示不限',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8 COMMENT='友情链接表';

DROP TABLE IF EXISTS `ll_log`;
CREATE TABLE `ll_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '网站日志ID 主键',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `code` varchar(64) NOT NULL DEFAULT '' COMMENT '日志代码',
  `route` varchar(64) NOT NULL DEFAULT '' COMMENT '请求路由',
  `url` varchar(2048) NOT NULL DEFAULT '' COMMENT '请求地址',
  `client_ip` varchar(39) NOT NULL DEFAULT '' COMMENT '客户端IP',
  `method` varchar(32) NOT NULL DEFAULT '' COMMENT '请求方式',
  `request` varchar(2048) NOT NULL DEFAULT '' COMMENT '请求参数',
  `content` varchar(2048) NOT NULL,
  `level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '等级：1=普通日志；2=警告日志；3=错误日志',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_userid` (`user_id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='网站日志表';

DROP TABLE IF EXISTS `ll_menu`;
CREATE TABLE `ll_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID 主键',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '资源名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级资源ID',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序 升序排列',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '菜单状态：1=启用；0=关闭',
  `icon` varchar(64) NOT NULL DEFAULT '' COMMENT '顶级菜单图标',
  `menu_route` varchar(100) NOT NULL DEFAULT '' COMMENT '相对admin模块的路由  #表示不是链接',
  `routes` varchar(512) NOT NULL DEFAULT '' COMMENT '此菜单下，授予多个路由，则每行写一个路由，如：\r\nindex/main\r\nindex/login',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '资源描述',
  `create_time` datetime NOT NULL DEFAULT now() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理菜单信息表';

DROP TABLE IF EXISTS `ll_sms`;
CREATE TABLE `ll_sms` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '短信记录ID 主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '短信类型：0=未指定；1=注册；2=登录；3=找回密码；4=通知；5=实名认证',
  `mobile` varchar(32) NOT NULL DEFAULT '' COMMENT '用户存储emai或者mobile',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '发送状态：0=未发送；1=已发送；2=发送失败',
  `msg` varchar(255) NOT NULL DEFAULT '' COMMENT '接收信息',
  `msg_no` varchar(32) NOT NULL DEFAULT '' COMMENT '消息编号',
  `return_code` varchar(32) NOT NULL DEFAULT '' COMMENT '服务商返回的状态代码',
  `return_info` varchar(255) NOT NULL DEFAULT '' COMMENT '短信服务商返回的全部信息',
  `client_ip` varchar(39) NOT NULL DEFAULT '' COMMENT '客户端IP',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '发送短信请求的URL',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='短信发送表';

DROP TABLE IF EXISTS `ll_cron`;
CREATE TABLE `ll_cron` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '计划任务ID 主键',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '任务名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '执行的URL,绝对地址',
  `year_start` int(11) NOT NULL DEFAULT '-1' COMMENT '开始年份：正数表示开始的年份，如2015表示2015年开始执行；-1表示不限制',
  `year_end` int(11) NOT NULL DEFAULT '-1' COMMENT '结束年份：正数表示结束年份，如2015表示到2015年结束；-1表示不限制',
  `month_start` int(11) NOT NULL DEFAULT '-1' COMMENT '开始月份：正数表示开始的月份，如3表示3月开始执行；-1表示不限制',
  `month_end` int(11) NOT NULL DEFAULT '-1' COMMENT '结束月份：正数表示结束的月份，如3表示3月结束执行；-1表示不限制',
  `day_start` int(11) NOT NULL DEFAULT '-1' COMMENT '星期几开始：如3表示星期三开始执行；-1表示不限制',
  `day_end` int(11) NOT NULL DEFAULT '-1' COMMENT '星期几结束：如3表示星期三结束执行；-1表示不限制',
  `date_start` int(11) NOT NULL DEFAULT '-1' COMMENT '开始天：如3表示3号开始执行；-1表示不限制',
  `date_end` int(11) NOT NULL DEFAULT '-1' COMMENT '结束的天：如3表示3号结束执行；-1表示不限制',
  `hour_start` int(11) NOT NULL DEFAULT '-1' COMMENT '开始小时：24小时制，如16表示16点开始执行；-1表示不限制',
  `hour_end` int(11) NOT NULL DEFAULT '-1' COMMENT '结束小时：24小时制，如15表示15点结束执行；-1表示不限制',
  `minute_start` int(11) NOT NULL DEFAULT '-1' COMMENT '开始的分钟：如15表示15分开始执行；-1表示不限制',
  `minute_end` int(11) NOT NULL DEFAULT '-1' COMMENT '结束的分钟：如15表示15分结束执行；-1表示不限制',
  `task_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '上次执行时间',
  `sort` int(11) NOT NULL DEFAULT '100' COMMENT '排序，升序执行',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='计划任务表';

DROP TABLE IF EXISTS `ll_cron_log`;
CREATE TABLE `ll_cron_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '计划任务日志ID 主键',
  `cron_id` int(11) NOT NULL DEFAULT '0' COMMENT '计划任务ID',
  `content` varchar(2480) NOT NULL DEFAULT '' COMMENT '执行的结果',
  `create_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='计划任务执行日志表';
