INSERT INTO tickers (TickerSymbol, type, lastupdated, updatefrequency) VALUES
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

-- these ETFs track commodities
INSERT INTO tickers (TickerSymbol, type, lastupdated, updatefrequency) VALUES
('PALL', 'commodity ETF', current_date - interval '4 months', 'm'), -- Palladium
('GLD', 'commodity ETF', current_date - interval '4 months', 'm'),
('SLV', 'commodity ETF', current_date - interval '4 months', 'm'), -- silver
('PPLT', 'commodity ETF', current_date - interval '4 months', 'm'), -- Platinum
('JJNTF', 'commodity ETF', current_date - interval '4 months', 'm'), -- Nickel
('JJUFF', 'commodity ETF', current_date - interval '4 months', 'm'), -- Aluminium
('JJCTF', 'commodity ETF', current_date - interval '4 months', 'm'), -- Copper
('TIN', 'commodity ETF', current_date - interval '4 months', 'm'), -- JJTFF
('REMX', 'commodity ETF', current_date - interval '4 months', 'm'), -- rare earth/strat metals

('JO', 'commodity ETF', current_date - interval '4 months', 'm'), -- coffee
('CANE', 'commodity ETF', current_date - interval '4 months', 'm'), -- sugar
('NIB', 'commodity ETF', current_date - interval '4 months', 'm'), -- Coca
-- ('CORN', 'commodity ETF', current_date - interval '4 months', 'm'), -- corn conflicts with alphavantage func corn
('SOYB', 'commodity ETF', current_date - interval '4 months', 'm'), --
('WEAT', 'commodity ETF', current_date - interval '4 months', 'm'),

('USO', 'commodity ETF', current_date - interval '4 months', 'm'), -- US oil
('BNO', 'commodity ETF', current_date - interval '4 months', 'm'), -- Brent
('UNG', 'commodity ETF', current_date - interval '4 months', 'm'), -- Nat Gas
('UGA', 'commodity ETF', current_date - interval '4 months', 'm'), -- Gasoline
('NLR', 'commodity ETF', current_date - interval '4 months', 'm'), -- nuclear
('PXE', 'commodity ETF', current_date - interval '4 months', 'm'), -- O&G exploration
('FCG', 'commodity ETF', current_date - interval '4 months', 'm'), -- nat gas producers
('URA', 'commodity ETF', current_date - interval '4 months', 'm'), -- Uranium

('TAN', 'commodity ETF', current_date - interval '4 months', 'm'), -- solar
('FAN', 'commodity ETF', current_date - interval '4 months', 'm'), -- wind
('PBD', 'commodity ETF', current_date - interval '4 months', 'm'), -- clean energy
('PIO', 'commodity ETF', current_date - interval '4 months', 'm'), -- global water (purification conversation) companies

('PPA', 'sector ETF', current_date - interval '4 months', 'm'), -- aerospace defense
('EVX', 'sector ETF', current_date - interval '4 months', 'm'), -- enviromental services
('ENFR', 'sector ETF', current_date - interval '4 months', 'm'), -- infra
('BJK', 'sector ETF', current_date - interval '4 months', 'm'), -- gaming
('PBJ', 'sector ETF', current_date - interval '4 months', 'm'), -- food and beverage consumer stapple
('XRT', 'sector ETF', current_date - interval '4 months', 'm'), -- retail
('IBUY', 'sector ETF', current_date - interval '4 months', 'm'), -- online retail
('PEJ', 'sector ETF', current_date - interval '4 months', 'm'), -- leisure and entertainment consumer discretionary
('GGME', 'sector ETF', current_date - interval '4 months', 'm'), -- media
('VICE', 'sector ETF', current_date - interval '4 months', 'm'), -- vice

('IAI', 'financial ETF', current_date - interval '4 months', 'm'), -- broker dealers
('PSP', 'financial ETF', current_date - interval '4 months', 'm'), -- private equity
('KBWP', 'financial ETF', current_date - interval '4 months', 'm'), -- property insurance
('KCE', 'financial ETF', current_date - interval '4 months', 'm'), -- capital markets
('KIE', 'financial ETF', current_date - interval '4 months', 'm'), -- insurance
('KBWB', 'financial ETF', current_date - interval '4 months', 'm'), -- banks 
('KBWR', 'financial ETF', current_date - interval '4 months', 'm'), -- regional banks

('IYT', 'sector ETF', current_date - interval '4 months', 'm'), -- transport
('GII', 'sector ETF', current_date - interval '4 months', 'm'), -- global infra
('SEA', 'sector ETF', current_date - interval '4 months', 'm'), -- shipping
('JETS', 'sector ETF', current_date - interval '4 months', 'm'), -- airlines
('CARZ', 'sector ETF', current_date - interval '4 months', 'm'), -- auto makers

('GDX', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- gold miners
('HAP', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- agro energy metals
('GDXJ', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- junior gold miner
('XME', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- metals and mining
('SLX', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- steel producers global
('SIL', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- silver miners
('WOOD', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- wood
('COPX', 'sector commodity ETF', current_date - interval '4 months', 'm'), -- copper producers

('PSI', 'commodity ETF', current_date - interval '4 months', 'm'), -- semiconductors
('KNCT', 'commodity ETF', current_date - interval '4 months', 'm'), -- networking ?
('IGPT', 'commodity ETF', current_date - interval '4 months', 'm'), -- software
('PNQI', 'commodity ETF', current_date - interval '4 months', 'm'), -- internet
('SKYY', 'commodity ETF', current_date - interval '4 months', 'm'), -- cloud computing
('NXTG', 'commodity ETF', current_date - interval '4 months', 'm'), -- smartphones
('SOCL', 'commodity ETF', current_date - interval '4 months', 'm'), -- social media
('ROBO', 'commodity ETF', current_date - interval '4 months', 'm'), -- robotics
('BLOK', 'commodity ETF', current_date - interval '4 months', 'm'), -- blockchain
('AIQ', 'commodity ETF', current_date - interval '4 months', 'm'), -- big data
('CIBR', 'commodity ETF', current_date - interval '4 months', 'm'), -- cybersecurity
--reits
-- global
('VNQI', 'strategy ETF', current_date - interval '4 months', 'q'), -- World ex-U.S.
('IFGL', 'sector ETF', current_date - interval '4 months', 'q'), -- Developed (ex-U.S.)
('VNQ', 'sector ETF', current_date - interval '4 months', 'q'), -- US
('RWO', 'sector ETF', current_date - interval '4 months', 'q'), -- World
-- U.S. Sector REITs
('INDS', 'sector ETF', current_date - interval '4 months', 'q'), -- Industrial
('REM', 'sector ETF', current_date - interval '4 months', 'q'), -- Mortgage
('REZ', 'sector ETF', current_date - interval '4 months', 'q'), -- Residential
('NETL', 'sector ETF', current_date - interval '4 months', 'q'), -- Net-Lease

('XHB', 'sector ETF', current_date - interval '4 months', 'q'), -- Homebuilders
('ROOF', 'sector ETF', current_date - interval '4 months', 'q'), -- Small Cap REITs
('HOMZ', 'sector ETF', current_date - interval '4 months', 'q'), -- Total U.S. Housing Market
('NURE', 'sector ETF', current_date - interval '4 months', 'q'), -- Short-Term REITs
-- baskets:
('DBB', 'commodity ETF', current_date - interval '4 months', 'm'), -- industrial metals
('BCI', 'commodity ETF', current_date - interval '4 months', 'm'), -- Bloomberg divers. commod.
('DBA', 'commodity ETF', current_date - interval '4 months', 'm'), -- agro
('DBE', 'commodity ETF', current_date - interval '4 months', 'm'), -- energy (holds futures in brent, coal, gas etc.)
('GSG', 'commodity ETF', current_date - interval '4 months', 'm') -- Goldman divers. commod.
;
-- these ETFs track themes in emerging markets
INSERT INTO tickers (TickerSymbol, type, lastupdated, updatefrequency) VALUES
('EMQQ', 'emerging market ETF', current_date - interval '4 months', 'q'), -- EM tech
('ECON', 'emerging market ETF', current_date - interval '4 months', 'q'), -- EM consumer
('EMIF', 'emerging market ETF', current_date - interval '4 months', 'q'), -- EM infrastructure

('CHIE', 'Commodity China ETF', current_date - interval '4 months', 'q'), -- China Energy
('CHIM', 'Commodity China ETF', current_date - interval '4 months', 'q'), -- China materials
('CHII', 'Industrial China ETF', current_date - interval '4 months', 'q'), -- China industrials
('CHIX', 'financial China ETF', current_date - interval '4 months', 'q'), -- China financials
('CHIR', 'China ETF', current_date - interval '4 months', 'q'), -- China realestate
('KURE', 'China ETF', current_date - interval '4 months', 'q'), -- China healthcare
('KWEB', 'China ETF', current_date - interval '4 months', 'q'), -- China internet
('CHIQ', 'China ETF', current_date - interval '4 months', 'q'), -- China consumer
('CQQQ', 'China ETF', current_date - interval '4 months', 'q'), -- China tech
('KFYP', 'China ETF', current_date - interval '4 months', 'q'), -- China infrastructure

('ECNS', 'China ETF', current_date - interval '4 months', 'q'), -- China small cap
('EWZS', 'ETF', current_date - interval '4 months', 'q'), -- Brazil small cap
('BRF', 'ETF', current_date - interval '4 months', 'q'), -- Brazil mid cap
('SMIN', 'ETF', current_date - interval '4 months', 'q'), -- India mid cap

('INCO', 'commodity ETF', current_date - interval '4 months', 'q'), -- India consumer
-- regions:
('BKF', 'regional ETF', current_date - interval '4 months', 'q'), -- Br, Ru, Ind, Ch
('IEMG', 'regional ETF', current_date - interval '4 months', 'q'), -- EM
('GMF', 'regional ETF', current_date - interval '4 months', 'q'), -- EM Asia-Pacific
('EMXC', 'regional ETF', current_date - interval '4 months', 'q'), -- EM ex-China
('EWX', 'regional ETF', current_date - interval '4 months', 'q'), -- EM Small caps
('AFK', 'regional ETF', current_date - interval '4 months', 'q'), -- Africa
('FM', 'regional ETF', current_date - interval '4 months', 'q'), -- Frontier countries ??
('ILF', 'regional ETF', current_date - interval '4 months', 'q'), -- Lat AM
('ASEA', 'regional ETF', current_date - interval '4 months', 'q') -- ASEAN
;

-- ETFs tracking sectors
INSERT INTO tickers (TickerSymbol, type, lastupdated, updatefrequency) VALUES
('XLB', 'sector commodity ETF', current_date - interval '4 months', 'q'), -- Basic Materials
('XLC', 'sector ETF', current_date - interval '4 months', 'q'), -- communication
('XLY', 'sector ETF', current_date - interval '4 months', 'q'), -- consummer discretionary
('XLP', 'sector ETF', current_date - interval '4 months', 'q'), -- consumer staples
('XLE', 'commodity sector ETF', current_date - interval '4 months', 'q'), -- energy companies
('XLF', 'financial sector ETF', current_date - interval '4 months', 'q'), --  financial 
('XLV', 'sector ETF', current_date - interval '4 months', 'q'), -- healthcare
('XLI', 'sector ETF', current_date - interval '4 months', 'q'), -- industrial
('XLK', 'sector ETF', current_date - interval '4 months', 'q'), --  tech
('XLU', 'sector ETF', current_date - interval '4 months', 'q'), -- utilities
('XLRE', 'sector ETF', current_date - interval '4 months', 'q'), -- real estate
-- US small cap sectors:
('PSCD', 'sector PSCD', current_date - interval '4 months', 'q'), -- small consumer discretionary
('PSCC', 'sector ETF', current_date - interval '4 months', 'q'), -- small consumer staples
('PSCH', 'sector ETF', current_date - interval '4 months', 'q'), --  healthcare
('PSCF', 'financial sector ETF', current_date - interval '4 months', 'q'), -- financial
('PSCE', 'commodity sector ETF', current_date - interval '4 months', 'q'), -- energy
('PSCI', 'sector ETF', current_date - interval '4 months', 'q'), -- industrial
('PSCM', 'commodity sector ETF', current_date - interval '4 months', 'q'), -- basic materials
('PSCU', 'sector ETF', current_date - interval '4 months', 'q'), -- utilities
('PSCT', 'sector ETF', current_date - interval '4 months', 'q'), -- tech
-- global sectors:
('MXI', 'commodity sector ETF', current_date - interval '4 months', 'q'), -- basic materials
('RXI', 'sector ETF', current_date - interval '4 months', 'q'), -- consumer discretionary
('KXI', 'sector ETF', current_date - interval '4 months', 'q'), -- consumer stapples
('IXC', 'commodity sector ETF', current_date - interval '4 months', 'q'), -- energy
('IXG', 'financial sector ETF', current_date - interval '4 months', 'q'), -- financial
('EXI', 'sector ETF', current_date - interval '4 months', 'q'), -- industrial
('IXN', 'sector ETF', current_date - interval '4 months', 'q'), -- tech
('IXP', 'sector ETF', current_date - interval '4 months', 'q'), -- telecom
('JXI', 'sector ETF', current_date - interval '4 months', 'q') -- utilities
;

-- ETFs tracking global things - think of better name and types, low discoverability
INSERT INTO tickers (TickerSymbol, type, lastupdated, updatefrequency) VALUES
-- quality
('IQLT', 'ETF', current_date - interval '4 months', 'm'), -- developped markets quality
('GEM', 'emerging market ETF', current_date - interval '4 months', 'm'), -- EM quality
('QWLD',  'ETF', current_date - interval '4 months', 'm'), -- global quality
('IQDF', 'ETF', current_date - interval '4 months', 'm'), -- global dividends
('SPHQ', 'sector ETF', current_date - interval '4 months', 'm'), -- quality large cap
('QUAL', 'sector ETF', current_date - interval '4 months', 'm') -- quality broad
;

-- still more etfs
INSERT INTO tickers (TickerSymbol, type, lastupdated, updatefrequency) VALUES
-- strategies
('PKW', 'strategy ETF', current_date - interval '4 months', 'q'), -- Buybacks
('PBP', 'strategy ETF', current_date - interval '4 months', 'q'), -- Buy-Write
('CEFS', 'strategy ETF', current_date - interval '4 months', 'q'), -- Closed-End Funds
('MOAT', 'strategy ETF', current_date - interval '4 months', 'q'), -- Wide Moat
('EQL', 'strategy ETF', current_date - interval '4 months', 'q'), -- Equal Weight Sectors
('QAI', 'strategy ETF', current_date - interval '4 months', 'q'), -- Hedge Multi-Strategy
('MNA', 'strategy ETF', current_date - interval '4 months', 'q'), -- Merger Arbitrage
('CSD', 'strategy ETF', current_date - interval '4 months', 'q'), -- Spin-Offs
('RYJ', 'strategy ETF', current_date - interval '4 months', 'q'), -- Strong Buy Rated Stocks
-- ('CPI', 'strategy ETF', current_date - interval '4 months', 'q'), -- Real Return #todo conflicts with...
('PIZ', 'strategy ETF', current_date - interval '4 months', 'q'), -- Technicals: Developed Markets
('PIE', 'strategy ETF', current_date - interval '4 months', 'q'), -- Technicals: Emerging Markets
('FPX', 'strategy ETF', current_date - interval '4 months', 'q'), -- U.S. IPOs
('SPHB', 'strategy ETF', current_date - interval '4 months', 'q'), -- High Beta
('SPLV', 'strategy ETF', current_date - interval '4 months', 'q'), -- Low Volatility
('IPKW', 'strategy ETF', current_date - interval '4 months', 'q'), -- International Buybacks
('GURU', 'strategy ETF', current_date - interval '4 months', 'q'), -- Guru Picks
('PHDG', 'strategy ETF', current_date - interval '4 months', 'q'), -- Hedged Equity
-- alternative asset classes
('HEDJ', 'strategy ETF', current_date - interval '4 months', 'q'), -- Int. Hedged Equity
('VXX', 'strategy ETF', current_date - interval '4 months', 'q'), -- Volatility - short term
('VXZ', 'strategy ETF', current_date - interval '4 months', 'q'), -- Volatility - mid-term
-- mixed fundementals
('PXF', 'strategy ETF', current_date - interval '4 months', 'q'), -- Developed Markets ex-U.S.
('PDN', 'strategy ETF', current_date - interval '4 months', 'q'), -- Developed Markets ex-U.S., small & mid caps
('PXH', 'emerging markets strategy ETF', current_date - interval '4 months', 'q'), -- Emerging Markets
('PRF', 'strategy ETF', current_date - interval '4 months', 'q'), -- U.S. 1000
('PRFZ', 'strategy ETF', current_date - interval '4 months', 'q'), -- U.S. 1500 Small & Mid Caps
-- by earnings
('EPS', 'strategy ETF', current_date - interval '4 months', 'q'), -- U.S. Large Cap
('EZM', 'strategy ETF', current_date - interval '4 months', 'q'), -- U.S. Mid Cap
('EES', 'strategy ETF', current_date - interval '4 months', 'q'), -- U.S. Small Cap
-- ESG
('DSI', 'strategy ETF', current_date - interval '4 months', 'q'), -- Social
('SUSL', 'strategy ETF', current_date - interval '4 months', 'q'), -- ESG - U.S.
('ESGG', 'strategy ETF', current_date - interval '4 months', 'q'), -- ESG - Global
('SHE', 'strategy ETF', current_date - interval '4 months', 'q'), -- Women in Leadership
('ETHO', 'strategy ETF', current_date - interval '4 months', 'q'), -- Environmental
-- asset allocation
('AOA', 'strategy ETF', current_date - interval '4 months', 'q'), -- Aggressive Asset Allocation
('AOK', 'strategy ETF', current_date - interval '4 months', 'q'), -- Conservative Asset Allocation
('AOR', 'strategy ETF', current_date - interval '4 months', 'q'), -- Growth Asset Allocation
('AOM', 'strategy ETF', current_date - interval '4 months', 'q') -- Moderate Asset Allocation
;

-- bond ETFs
INSERT INTO tickers (TickerSymbol, type, lastupdated, updatefrequency) VALUES
('US1M', 'bond ETF', current_date - interval '4 months', 'q'), -- 1 Month Treasury
('US2M', 'bond ETF', current_date - interval '4 months', 'q'), -- 2 Month Treasury
('US3M', 'bond ETF', current_date - interval '4 months', 'q'), -- 3 Month Treasury
('US6M', 'bond ETF', current_date - interval '4 months', 'q'), -- 6 Month Treasury
('US12M', 'bond ETF', current_date - interval '4 months', 'q'), -- 12 Month Treasury
('US2Y', 'bond ETF', current_date - interval '4 months', 'q'), -- 2 Year Treasury
('US5Y', 'bond ETF', current_date - interval '4 months', 'q'), -- 5 Year Treasury
('US7Y', 'bond ETF', current_date - interval '4 months', 'q'), -- 7 Year Treasury
('US10Y', 'bond ETF', current_date - interval '4 months', 'q'), -- 10 Year Treasury
('US20Y', 'bond ETF', current_date - interval '4 months', 'q'), -- 20 Year Treasury
('US30Y', 'bond ETF', current_date - interval '4 months', 'q'), -- 30 Year Treasury
('IGLB', 'bond ETF', current_date - interval '4 months', 'q'), -- U.S. Corporate Bonds - Long Term
('PFIG', 'bond ETF', current_date - interval '4 months', 'q'), -- Investment Grade Bonds - Broad
('IGIB', 'bond ETF', current_date - interval '4 months', 'q'), -- U.S. Corporate Bonds - Medium Term
('IGSB', 'bond ETF', current_date - interval '4 months', 'q'), -- Short Term Investment Grade Bonds
('PHB', 'bond ETF', current_date - interval '4 months', 'q'), -- High Yield Bonds - Broad
('PCY', 'emerging market bond ETF', current_date - interval '4 months', 'q'), -- Emerging Market Government Bonds
('WIP', 'bond ETF', current_date - interval '4 months', 'q'), -- Inflation Protected Int. Govt. Bonds
('BWX', 'bond ETF', current_date - interval '4 months', 'q'), -- International Govt. Bonds
('BWZ', 'bond ETF', current_date - interval '4 months', 'q'), -- Short Term Int. Govt. Bonds
('PICB', 'bond ETF', current_date - interval '4 months', 'q') -- International Corporate Bonds
;