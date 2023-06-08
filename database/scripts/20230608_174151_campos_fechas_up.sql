ALTER TABLE resoluciones_new.resolucion
    ADD COLUMN IF NOT EXISTS fecha_inicio DATE,
    ADD COLUMN IF NOT EXISTS fecha_fin DATE;
