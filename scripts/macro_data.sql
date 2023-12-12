INSERT INTO tickers (TickerSymbol, type, importance) VALUES
-- comodities
('WTI', 'commodity', 'm'),
('BRENT', 'commodity', 'm'),
('NATURAL_GAS', 'commodity', 'm'),
('COPPER', 'commodity', 'm'),
('ALUMINUM', 'commodity', 'm'),
('WHEAT', 'commodity', 'm'),
('CORN', 'commodity', 'm'),
('COTTON', 'commodity', 'm'),
('SUGAR', 'commodity', 'm'),
('ALL_COMMODITIES', 'commodity', 'm'),
-- macro indicators
('REAL_GDP', 'macro', 'm'),
('REAL_GDP_PER_CAPITA', 'macro', 'm'),
('FEDERAL_FUNDS_RATE', 'macro', 'm'),
('CPI', 'macro', 'm'),
('INFLATION', 'macro', 'm'),
('RETAIL_SALES', 'macro', 'm'),
('DURABLES', 'macro', 'm'),
('UNEMPLOYMENT', 'macro', 'm'),
('NONFARM_PAYROLL', 'macro', 'm'),
-- bond rates
('TREASURY_YIELD 3month', 'macro', 'm'),
('TREASURY_YIELD 2year', 'macro', 'm'),
('TREASURY_YIELD 5year', 'macro', 'm'),
('TREASURY_YIELD 7year', 'macro', 'm'),
('TREASURY_YIELD 10year', 'macro', 'm'),
('TREASURY_YIELD 30year', 'macro', 'm')
;