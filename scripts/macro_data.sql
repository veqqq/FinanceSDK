INSERT INTO tickers (TickerSymbol, type, lastupdated, importance) VALUES
-- comodities
('WTI', 'commodity', current_date - interval '4 months', 'm'),
('BRENT', 'commodity', current_date - interval '4 months', 'm'),
('NATURAL_GAS', 'commodity', current_date - interval '4 months', 'm'),
('COPPER', 'commodity', current_date - interval '4 months', 'm'),
('ALUMINUM', 'commodity', current_date - interval '4 months', 'm'),
('WHEAT', 'commodity', current_date - interval '4 months', 'm'),
('CORN', 'commodity', current_date - interval '4 months', 'm'),
('COTTON', 'commodity', current_date - interval '4 months', 'm'),
('SUGAR', 'commodity', current_date - interval '4 months', 'm'),
('ALL_COMMODITIES', 'commodity', current_date - interval '4 months', 'm'),
-- macro indicators
('REAL_GDP', 'macro', current_date - interval '4 months', 'm'),
('REAL_GDP_PER_CAPITA', 'macro', current_date - interval '4 months', 'm'),
('FEDERAL_FUNDS_RATE', 'macro', current_date - interval '4 months', 'm'),
('CPI', 'macro', current_date - interval '4 months', 'm'),
('INFLATION', 'macro', current_date - interval '4 months', 'm'),
('RETAIL_SALES', 'macro', current_date - interval '4 months', 'm'),
('DURABLES', 'macro', current_date - interval '4 months', 'm'),
('UNEMPLOYMENT', 'macro', current_date - interval '4 months', 'm'),
('NONFARM_PAYROLL', 'macro', current_date - interval '4 months', 'm'),
-- bond rates
('TREASURY_YIELD 3month', 'macro', current_date - interval '4 months', 'm'),
('TREASURY_YIELD 2year', 'macro', current_date - interval '4 months', 'm'),
('TREASURY_YIELD 5year', 'macro', current_date - interval '4 months', 'm'),
('TREASURY_YIELD 7year', 'macro', current_date - interval '4 months', 'm'),
('TREASURY_YIELD 10year', 'macro', current_date - interval '4 months', 'm'),
('TREASURY_YIELD 30year', 'macro', current_date - interval '4 months', 'm')
;