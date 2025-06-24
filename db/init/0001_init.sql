CREATE USER "raya-local" WITH PASSWORD 'raya-local';

CREATE DATABASE raya_local_corp
    OWNER "raya-local"
    ENCODING 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8';