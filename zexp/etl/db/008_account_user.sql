CREATE TABLE if not exists account_user (
    account_id bigint,
    user_id bigint,
    status int2,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    roles text[],
    permissions text[],
    full_name text,
    short_name text,
    position text,
    response_status int2,
    invitation_sent_at timestamp with time zone,
    invitation_sent_by int8,
    invitation_accepted_at timestamp with time zone,
    invitation_rejected_at timestamp with time zone,
    disabled_at timestamp with time zone,
    disabled_by int8,
    disable_reason text,
    rid int8,
    PRIMARY KEY (account_id, user_id)
);