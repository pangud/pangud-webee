CREATE TABLE ftp_groups (
                                 group_name varchar(16) NOT NULL default '',
                                 gid smallint(6) NOT NULL default '2001',
                                 members varchar(16) NOT NULL default '',
                                 KEY group_name (group_name)
) ENGINE=MyISAM COMMENT='ProFTP group table';

CREATE TABLE ftp_users (
                                id int(10) unsigned NOT NULL auto_increment,
                                username varchar(32) NOT NULL default '',
                                passwd varchar(32) NOT NULL default '',
                                uid smallint(6) NOT NULL default '1000',
                                gid smallint(6) NOT NULL default '1000',
                                homedir varchar(255) NOT NULL default '',
                                shell varchar(16) NOT NULL default '/sbin/nologin',
                                count int(11) NOT NULL default '0',
                                accessed TIMESTAMP NOT NULL default '2000-01-01 00:00:00',
                                modified TIMESTAMP NOT NULL default '2000-01-01 00:00:00',
                                PRIMARY KEY (id),
                                UNIQUE KEY userid (username)
) ENGINE=MyISAM COMMENT='ProFTP user table';

INSERT INTO `ftp_groups` (`group_name`, `gid`, `members`) VALUES ('ftp_group', 2001, 'ftp_user');
INSERT INTO `ftp_users` (`id`, `username`, `passwd`, `uid`, `gid`, `homedir`, `shell`, `count`, `accessed`, `modified`) VALUES (NULL, 'test', 'password', '1000', '1000', '/data/test', '/sbin/nologin', '0', '2000-01-01 00:00:00', '2000-01-01 00:00:00');
