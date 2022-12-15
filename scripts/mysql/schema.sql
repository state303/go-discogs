CREATE TABLE ` data `
(
    `etag` VARCHAR
(
    200
) UNIQUE PRIMARY KEY NOT NULL COMMENT '' ETag representing this data being unique.Used for updating cache.'',
    ` generated_at ` TIMESTAMP NOT NULL COMMENT '' Date this data is generated at.'',
    ` checksum ` VARCHAR
(
    200
) NOT NULL COMMENT '' Checksum to validate when fetching dump source.'',
    ` target_type ` VARCHAR
(
    20
) NOT NULL COMMENT '' Type of data.Referred as artists, labels, masters, release.
    Always uppercase.'',
    ` uri ` VARCHAR
(
    2048
) NOT NULL COMMENT '' URI to download dump data file.''
    );

CREATE TABLE ` style `
(
    `
    id
    `
    INTEGER
    PRIMARY
    KEY
    NOT
    NULL
    AUTO_INCREMENT
    COMMENT
    ''
    id
    of
    style
    '',
    `
    name
    `
    VARCHAR
(
    50
) UNIQUE NOT NULL COMMENT '' name of style ''
    );

CREATE TABLE ` genre `
(
    `
    id
    `
    INTEGER
    PRIMARY
    KEY
    NOT
    NULL
    AUTO_INCREMENT
    COMMENT
    ''
    id
    of
    genre
    '',
    `
    name
    `
    VARCHAR
(
    50
) UNIQUE NOT NULL COMMENT '' name of genre ''
    );

CREATE TABLE ` artist `
(
    `
    id
    `
    INTEGER
    PRIMARY
    KEY
    NOT
    NULL,
    `
    data_quality
    `
    VARCHAR
(
    100
),
    ` name ` VARCHAR
(
    1000
),
    ` profile ` TEXT,
    ` real_name ` VARCHAR
(
    2000
)
    );

CREATE TABLE ` label `
(
    `
    id
    `
    INTEGER
    PRIMARY
    KEY
    NOT
    NULL,
    `
    contact_info
    `
    TEXT,
    `
    data_quality
    `
    VARCHAR
(
    100
),
    ` name ` VARCHAR
(
    300
),
    ` profile ` TEXT,
    ` parent_id ` INTEGER
    );

CREATE TABLE ` master `
(
    `
    id
    `
    INTEGER
    PRIMARY
    KEY
    NOT
    NULL,
    `
    data_quality
    `
    VARCHAR
(
    100
),
    ` title ` VARCHAR
(
    2000
),
    ` released_year ` SMALLINT
    );

CREATE TABLE ` release `
(
    `
    id
    `
    INTEGER
    PRIMARY
    KEY
    NOT
    NULL,
    `
    title
    `
    VARCHAR
(
    10000
),
    ` country ` VARCHAR
(
    100
),
    ` data_quality ` VARCHAR
(
    100
),
    ` released_year ` SMALLINT,
    ` released_month ` SMALLINT,
    ` released_day ` SMALLINT,
    ` listed_release_date ` VARCHAR
(
    255
),
    ` master_id ` INTEGER NOT NULL COMMENT '' id of master release this release belongs to '',
    ` is_master ` BOOLEAN,
    ` notes ` TEXT,
    ` status ` VARCHAR
(
    255
)
    );

CREATE TABLE ` release_genre `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    genre_id
    `
    INTEGER
    NOT
    NULL,
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    genre_id
    `
)
    );

CREATE TABLE ` release_track `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    duration
    `
    VARCHAR
(
    1500
),
    ` position ` VARCHAR
(
    1500
),
    ` title ` VARCHAR
(
    10000
),
    ` title_hash ` BIGINT NOT NULL COMMENT '' fnv32 encoded hash from title '',
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    title_hash
    `
)
    );

CREATE TABLE ` label_release `
(
    `
    label_id
    `
    INTEGER
    NOT
    NULL,
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    category_notation
    `
    VARCHAR
(
    1000
),
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    label_id
    `,
    `
    release_id
    `
)
    );

CREATE TABLE ` release_image `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    url_hash
    `
    BIGINT
    NOT
    NULL
    COMMENT
    ''
    fnv32
    encoded
    hash
    from
    url
    '',
    `
    url
    `
    VARCHAR
(
    2048
) NOT NULL,
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    url_hash
    `
)
    );

CREATE TABLE ` release_contract `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    label_id
    `
    INTEGER
    NOT
    NULL,
    `
    contract_hash
    `
    BIGINT
    NOT
    NULL
    COMMENT
    ''
    fnv32
    encoded
    hash
    from
    contract
    '',
    `
    contract
    `
    VARCHAR
(
    5000
) NOT NULL,
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    label_id
    `,
    `
    contract_hash
    `
)
    );

CREATE TABLE ` release_identifier `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    description
    `
    TEXT,
    `
    type
    `
    VARCHAR
(
    2500
),
    ` value ` TEXT,
    ` identifier_hash ` BIGINT NOT NULL COMMENT '' fnv32 encoded hash from string which is description, type, value appended in order '',
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    identifier_hash
    `
)
    );

CREATE TABLE ` master_track `
(
    `
    master_id
    `
    INTEGER
    NOT
    NULL,
    `
    duration
    `
    VARCHAR
(
    1500
),
    ` position ` VARCHAR
(
    1500
),
    ` title ` VARCHAR
(
    10000
),
    ` title_hash ` BIGINT NOT NULL COMMENT '' fnv32 encoded hash from title '',
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    master_id
    `,
    `
    title_hash
    `
)
    );

CREATE TABLE ` master_video `
(
    `
    master_id
    `
    INTEGER
    NOT
    NULL,
    `
    url_hash
    `
    BIGINT
    NOT
    NULL
    COMMENT
    ''
    fnv32
    encoded
    hash
    from
    url
    '',
    `
    url
    `
    VARCHAR
(
    2048
) NOT NULL,
    ` description ` VARCHAR
(
    4000
),
    ` title ` VARCHAR
(
    1000
),
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    master_id
    `,
    `
    url_hash
    `
)
    );

CREATE TABLE ` master_genre `
(
    `
    master_id
    `
    INTEGER
    NOT
    NULL,
    `
    genre_id
    `
    INTEGER
    NOT
    NULL,
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    master_id
    `,
    `
    genre_id
    `
)
    );

CREATE TABLE ` master_style `
(
    `
    master_id
    `
    INTEGER
    NOT
    NULL,
    `
    style_id
    `
    INTEGER
    NOT
    NULL,
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    master_id
    `,
    `
    style_id
    `
)
    );

CREATE TABLE ` release_style `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    style_id
    `
    INTEGER
    NOT
    NULL,
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    style_id
    `
)
    );

CREATE TABLE ` release_video `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    description
    `
    VARCHAR
(
    4000
),
    ` title ` VARCHAR
(
    1000
),
    ` url ` VARCHAR
(
    2048
) NOT NULL,
    ` url_hash ` BIGINT NOT NULL COMMENT '' fnv32 encoded hash from url '',
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    url_hash
    `
)
    );

CREATE TABLE ` label_url `
(
    `
    label_id
    `
    INTEGER
    NOT
    NULL,
    `
    url_hash
    `
    BIGINT
    NOT
    NULL
    COMMENT
    ''
    fnv32
    encoded
    hash
    from
    url
    '',
    `
    url
    `
    VARCHAR
(
    2048
) NOT NULL,
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    label_id
    `,
    `
    url_hash
    `
)
    );

CREATE TABLE ` release_format `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    description
    `
    VARCHAR
(
    10000
),
    ` name ` VARCHAR
(
    255
),
    ` quantity ` INTEGER,
    ` text ` VARCHAR
(
    5000
),
    ` format_hash ` BIGINT NOT NULL COMMENT '' fnv32 encoded hash from string which is description, name, quantity, text appended in order '',
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    format_hash
    `
)
    );

CREATE TABLE ` artist_alias `
(
    `
    artist_id
    `
    INTEGER
    NOT
    NULL,
    `
    alias_id
    `
    INTEGER
    NOT
    NULL,
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    artist_id
    `,
    `
    alias_id
    `
)
    );

CREATE TABLE ` artist_name_variation `
(
    `
    artist_id
    `
    INTEGER
    NOT
    NULL
    COMMENT
    ''
    id
    of
    artist
    '',
    `
    name_variation
    `
    VARCHAR
(
    2000
) NOT NULL COMMENT '' artist\'' s other name '',
    ` name_variation_hash ` BIGINT NOT NULL COMMENT '' fnv32 encoded hash from name_variation '',
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    artist_id
    `,
    `
    name_variation_hash
    `
)
    );

CREATE TABLE ` master_artist `
(
    `
    artist_id
    `
    INTEGER
    NOT
    NULL
    COMMENT
    ''
    artist
    id
    of
    the
    master
    release
    '',
    `
    master_id
    `
    INTEGER
    NOT
    NULL
    COMMENT
    ''
    master
    id
    '',
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    artist_id
    `,
    `
    master_id
    `
)
    );

CREATE TABLE ` release_artist `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    artist_id
    `
    INTEGER
    NOT
    NULL,
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    artist_id
    `
)
    );

CREATE TABLE ` release_credited_artist `
(
    `
    release_id
    `
    INTEGER
    NOT
    NULL,
    `
    artist_id
    `
    INTEGER
    NOT
    NULL,
    `
    role_hash
    `
    BIGINT
    NOT
    NULL
    COMMENT
    ''
    fnv32
    encoded
    hash
    from
    role
    '',
    `
    role
    `
    VARCHAR
(
    10000
) COMMENT '' role of an artist for a release '',
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    release_id
    `,
    `
    artist_id
    `,
    `
    role_hash
    `
)
    );

CREATE TABLE ` artist_url `
(
    `
    artist_id
    `
    INTEGER
    NOT
    NULL,
    `
    url_hash
    `
    BIGINT
    NOT
    NULL
    COMMENT
    ''
    fnv32
    encoded
    hash
    from
    url
    '',
    `
    url
    `
    VARCHAR
(
    2048
) NOT NULL,
    ` updated_at ` TIMESTAMP NOT NULL DEFAULT
(
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    artist_id
    `,
    `
    url_hash
    `
)
    );

CREATE TABLE ` artist_group `
(
    `
    artist_id
    `
    INTEGER
    NOT
    NULL,
    `
    group_id
    `
    INTEGER
    NOT
    NULL,
    `
    updated_at
    `
    TIMESTAMP
    NOT
    NULL
    DEFAULT (
    NOW
(
)) COMMENT '' created time '',
    PRIMARY KEY
(
    `
    artist_id
    `,
    `
    group_id
    `
)
    );

CREATE INDEX ` pk_style ` ON ` style ` (` id `);

ALTER TABLE ` data ` COMMENT = '' Cached resource for keep tracking data dump updates (either being monthly or random occations)'';

ALTER TABLE ` label ` ADD CONSTRAINT ` fk_label_parent_id_label_id ` FOREIGN KEY (` parent_id `) REFERENCES ` label ` (` id `);

ALTER TABLE ` release_genre ` ADD CONSTRAINT ` fk_release_genre_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` release_genre ` ADD CONSTRAINT ` fk_release_genre_genre_id_genre ` FOREIGN KEY (` genre_id `) REFERENCES ` genre ` (` id `);

ALTER TABLE ` release_track ` ADD CONSTRAINT ` fk_release_track_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` label_release ` ADD CONSTRAINT ` fk_label_release_label_id_label ` FOREIGN KEY (` label_id `) REFERENCES ` label ` (` id `);

ALTER TABLE ` label_release ` ADD CONSTRAINT ` fk_label_release_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` release_image ` ADD CONSTRAINT ` fk_release_image_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` release_contract ` ADD CONSTRAINT ` fk_release_contract_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` release_contract ` ADD CONSTRAINT ` fk_release_contract_label_id_label ` FOREIGN KEY (` label_id `) REFERENCES ` label ` (` id `);

ALTER TABLE ` release_identifier ` ADD CONSTRAINT ` fk_release_identifier_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` master_video ` ADD CONSTRAINT ` fk_master_video_master_id_master ` FOREIGN KEY (` master_id `) REFERENCES ` master ` (` id `);

ALTER TABLE ` master_genre ` ADD CONSTRAINT ` fk_master_genre_master_id_master ` FOREIGN KEY (` master_id `) REFERENCES ` master ` (` id `);

ALTER TABLE ` master_genre ` ADD CONSTRAINT ` fk_master_genre_genre_id_genre ` FOREIGN KEY (` genre_id `) REFERENCES ` genre ` (` id `);

ALTER TABLE ` master_style ` ADD CONSTRAINT ` fk_master_style_master_id_master ` FOREIGN KEY (` master_id `) REFERENCES ` master ` (` id `);

ALTER TABLE ` master_style ` ADD CONSTRAINT ` fk_master_style_style_id_style ` FOREIGN KEY (` style_id `) REFERENCES ` style ` (` id `);

ALTER TABLE ` release_style ` ADD CONSTRAINT ` fk_release_style_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` release_style ` ADD CONSTRAINT ` fk_release_style_style_id_style ` FOREIGN KEY (` style_id `) REFERENCES ` style ` (` id `);

ALTER TABLE ` release_video ` ADD CONSTRAINT ` fk_release_video_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` label_url ` ADD CONSTRAINT ` fk_label_url_label_id_label ` FOREIGN KEY (` label_id `) REFERENCES ` label ` (` id `);

ALTER TABLE ` release_format ` ADD CONSTRAINT ` fk_release_format_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` artist_alias ` ADD CONSTRAINT ` fk_artist_alias_artist_id_artist ` FOREIGN KEY (` artist_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` artist_alias ` ADD CONSTRAINT ` fk_artist_alias_alias_id_artist ` FOREIGN KEY (` alias_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` artist_name_variation ` ADD CONSTRAINT ` fk_artist_name_variation_artist_id_artist ` FOREIGN KEY (` artist_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` master_artist ` ADD CONSTRAINT ` fk_master_artist_artist_id_artist ` FOREIGN KEY (` artist_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` master_artist ` ADD CONSTRAINT ` fk_master_artist_master_id_master ` FOREIGN KEY (` master_id `) REFERENCES ` master ` (` id `);

ALTER TABLE ` release_artist ` ADD CONSTRAINT ` fk_release_artist_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` release_artist ` ADD CONSTRAINT ` fk_release_artist_artist_id_artist ` FOREIGN KEY (` artist_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` release_credited_artist ` ADD CONSTRAINT ` fk_release_credited_artist_release_id_release ` FOREIGN KEY (` release_id `) REFERENCES ` release ` (` id `);

ALTER TABLE ` release_credited_artist ` ADD CONSTRAINT ` fk_release_credited_artist_artist_id_artist ` FOREIGN KEY (` artist_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` artist_url ` ADD CONSTRAINT ` fk_artist_url_artist_id_artist ` FOREIGN KEY (` artist_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` artist_group ` ADD CONSTRAINT ` fk_artist_group_artist_id_artist ` FOREIGN KEY (` artist_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` artist_group ` ADD CONSTRAINT ` fk_artist_group_group_id_artist ` FOREIGN KEY (` group_id `) REFERENCES ` artist ` (` id `);

ALTER TABLE ` release ` ADD CONSTRAINT ` fk_release_master_id_master ` FOREIGN KEY (` master_id `) REFERENCES ` master ` (` id `);

