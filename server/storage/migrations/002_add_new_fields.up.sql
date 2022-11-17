ALTER TABLE resources ADD COLUMN language text;
ALTER TABLE resources ADD COLUMN type text;

ALTER TABLE travelers ADD COLUMN image text;
ALTER TABLE travelers ADD COLUMN nationality text;

ALTER TABLE travels ADD COLUMN route text;