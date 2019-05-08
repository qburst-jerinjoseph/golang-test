BEGIN;

INSERT INTO sample (
  text
) VALUES
    ('lazygo')
ON CONFLICT DO NOTHING;

COMMIT;