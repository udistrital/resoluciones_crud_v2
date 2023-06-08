ALTER TABLE resoluciones_new.resoluciones
    ADD COLUMN IF NOT EXISTS fecha_inicio DATE,
    ADD COLUMN IF NOT EXISTS fecha_fin DATE;
