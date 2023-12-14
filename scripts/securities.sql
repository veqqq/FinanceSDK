insert into tickers (tickersymbol, lastupdated, importance, type) values
    -- country etfs
    ('EWZ', current_date - interval '4 months', 'm', 'Country ETF'), -- Brazil
    ('EWC', current_date - interval '4 months', 'q', 'Country ETF'), -- Cannada
    ('GXC', current_date - interval '4 months', 'q', 'Country ETF'), -- China
    ('EWG', current_date - interval '4 months', 'q', 'Country ETF'), -- Germany
    ('EWQ', current_date - interval '4 months', 'q', 'Country ETF'), -- France
    ('EWH', current_date - interval '4 months', 'q', 'Country ETF'), -- Hong Kong
    ('ECH', current_date - interval '4 months', 'q', 'Country ETF'), -- Chile
    ('EIRL', current_date - interval '4 months', 'q', 'Country ETF'), -- Ireland
    ('EIS', current_date - interval '4 months', 'q', 'Country ETF'), -- Israel

    ('EWI', current_date - interval '4 months', 'q', 'Country ETF'), -- Italy
    ('IDX', current_date - interval '4 months', 'q', 'Country ETF'), -- Indonesia
    ('EWM', current_date - interval '4 months', 'q', 'Country ETF'), -- Malaysia
    ('EWW', current_date - interval '4 months', 'q', 'Country ETF'), -- Mexico
    ('EWN', current_date - interval '4 months', 'q', 'Country ETF'), -- Netherlands
    ('EWS', current_date - interval '4 months', 'q', 'Country ETF'), -- Singapore
    ('EWY', current_date - interval '4 months', 'q', 'Country ETF'), -- S. Korea
    ('EWL', current_date - interval '4 months', 'q', 'Country ETF'), -- Switzerland
    ('EWD', current_date - interval '4 months', 'q', 'Country ETF'), -- Sweden
    ('THD', current_date - interval '4 months', 'q', 'Country ETF'), -- Thailand
    ('TUR', current_date - interval '4 months', 'q', 'Country ETF'), -- 
    
    ('EZA', current_date - interval '4 months', 'q', 'Country ETF'), -- S. Africa
    ('EWP', current_date - interval '4 months', 'q', 'Country ETF'), -- Spain
    ('EWU', current_date - interval '4 months', 'q', 'Country ETF'), -- U.K.
    ('VNM', current_date - interval '4 months', 'q', 'Country ETF'), -- Vietnam
    ('NORW', current_date - interval '4 months', 'q', 'Country ETF'), -- Norway
    ('EPHE', current_date - interval '4 months', 'q', 'Country ETF'), -- Philippines
    -- ('NGE', current_date - interval '4 months', 'q', 'Country ETF'), -- Nigeria not supported on alphavantage
    ('QAT', current_date - interval '4 months', 'q', 'Country ETF'), -- Qatar
    ('UAE', current_date - interval '4 months', 'q', 'Country ETF'), -- 
    ('EWJ', current_date - interval '4 months', 'q', 'Country ETF'), -- Japan
    ('KSA', current_date - interval '4 months', 'q', 'Country ETF'), -- Saudi Arabia
    ('ARGT', current_date - interval '4 months', 'q', 'Country ETF'), -- Argentina
    ('KWT', current_date - interval '4 months', 'q', 'Country ETF'), -- 
    ('EFNL', current_date - interval '4 months', 'q', 'Country ETF') -- Finland
;


insert into datasources (sourceid, sourcename, sourceurl) values 
(1, 'Alphavantage', 'https://www.alphavantage.co and https://alpha-vantage.p.rapidapi.com');


-- weird tickers, some like FUCC for Formosan Union Chemical Corp don't work?
-- EAT, PLAY, CAKE, KOOL, MMM, MOO, FUN, PZZA
-- CUM ASS BJ NUT GAY FUCC OUCH TWNK BBW HEINY BUNZ FUC OMG RACE BOOM LUV WTH
-- ZZ, BABY, GRR,, COOL, 


-- stocks I like
insert into tickers (tickersymbol, lastupdated, importance) values
    -- energy
    ('LEU', 'commodity stock', current_date - interval '4 months', 'm'), -- nuclear stuff uranium
    -- banks
    ('WAL', 'financial stock', current_date - interval '4 months', 'm'),
    ('BDORY', 'financial stock', current_date - interval '4 months', 'm')
;

insert into tickers (tickersymbol, type, lastupdated, importance) values
    -- steel
    ('CLF', 'commodity stock', current_date - interval '4 months', 'm'),
    ('MT', 'commodity stock', current_date - interval '4 months', 'm'),
    ('X', 'commodity stock', current_date - interval '4 months', 'm'),
    ('STLD', 'commodity stock', current_date - interval '4 months', 'm'),
    ('GGB', 'commodity stock', current_date - interval '4 months', 'm'),
    ('TX', 'commodity stock', current_date - interval '4 months', 'm'),
    ('RS', 'commodity stock', current_date - interval '4 months', 'm'),
    ('ASTL', 'commodity stock', current_date - interval '4 months', 'm'),
    ('IIN', 'commodity stock', current_date - interval '4 months', 'm'),
    ('SPLP', 'commodity stock', current_date - interval '4 months', 'm'),
    ('STCN', 'commodity stock', current_date - interval '4 months', 'm'),
    ('WS', 'commodity stock', current_date - interval '4 months', 'm'),
    -- coal
    ('HCC', 'commodity stock', current_date - interval '4 months', 'm'),
    ('AMR', 'commodity stock', current_date - interval '4 months', 'm'),
    ('BTU', 'commodity stock', current_date - interval '4 months', 'm'),
    ('ARCH', 'commodity stock', current_date - interval '4 months', 'm'),
    ('NRP', 'commodity stock', current_date - interval '4 months', 'm'),
    ('METC', 'commodity stock', current_date - interval '4 months', 'm'),
    ('CEIX', 'commodity stock', current_date - interval '4 months', 'm'),
    ('HNRG', 'commodity stock', current_date - interval '4 months', 'm'),
    ('SXC', 'commodity stock', current_date - interval '4 months', 'm'),
    ('ARLP', 'commodity stock', current_date - interval '4 months', 'm'),
    ('AREC', 'commodity stock', current_date - interval '4 months', 'm'),
    ('NC', 'commodity stock', current_date - interval '4 months', 'm'),
    -- other commodity
    ('TECK', 'commodity stock', current_date - interval '4 months', 'm') -- sold coal to glencore
    ;

    -- energy oil
    insert into tickers (tickersymbol, type, lastupdated, importance) values
('PAGP', 'Buffet oil commodity stock', current_date - interval '4 months', 'm'),
('PBF', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PBFX', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PDCE', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PSX', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PTEN', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PUMP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('RES', 'oil commodity stock', current_date - interval '4 months', 'm'),
('RIG', 'oil commodity stock', current_date - interval '4 months', 'm'),
('RRC', 'oil commodity stock', current_date - interval '4 months', 'm'),
('SLB', 'oil commodity stock', current_date - interval '4 months', 'm'),
('SM', 'oil commodity stock', current_date - interval '4 months', 'm'),
('SPH', 'oil commodity stock', current_date - interval '4 months', 'm'),
('SUN', 'oil commodity stock', current_date - interval '4 months', 'm'),
('SWN', 'oil commodity stock', current_date - interval '4 months', 'm'),
('TRGP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('UGP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('VLO', 'oil commodity stock', current_date - interval '4 months', 'm'),
('VNOM', 'oil commodity stock', current_date - interval '4 months', 'm'),
('WES', 'oil commodity stock', current_date - interval '4 months', 'm'),
('WHD', 'oil commodity stock', current_date - interval '4 months', 'm'),
('WMB', 'oil commodity stock', current_date - interval '4 months', 'm'),
('XEC', 'oil commodity stock', current_date - interval '4 months', 'm'),
('XOM', 'oil commodity stock', current_date - interval '4 months', 'm'),
('YPF', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PBR', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PBR-A', 'oil commodity stock', current_date - interval '4 months', 'm'),
('MRO', 'oil commodity stock', current_date - interval '4 months', 'm'),
('OXY', 'Buffet oil commodity stock', current_date - interval '4 months', 'm'),
('SHEL', 'oil commodity stock', current_date - interval '4 months', 'm'),
('SU', 'oil commodity stock', current_date - interval '4 months', 'm'),
('FECCF', 'oil commodity stock', current_date - interval '4 months', 'm'),
('TTE', 'oil commodity stock', current_date - interval '4 months', 'm'),
('EQNR', 'oil commodity stock', current_date - interval '4 months', 'm'),
('CNQ', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PXD', 'oil commodity stock', current_date - interval '4 months', 'm'),
('TRP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('OKE', 'oil commodity stock', current_date - interval '4 months', 'm'),
('WDS', 'oil commodity stock', current_date - interval '4 months', 'm'),
('CRK', 'oil commodity stock', current_date - interval '4 months', 'm'),
('CVX', 'oil commodity stock', current_date - interval '4 months', 'm'),
('DCP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('DK', 'oil commodity stock', current_date - interval '4 months', 'm'),
('DRQ', 'oil commodity stock', current_date - interval '4 months', 'm'),
('DVN', 'oil commodity stock', current_date - interval '4 months', 'm'),
('E', 'oil commodity stock', current_date - interval '4 months', 'm'),
('ENB', 'oil commodity stock', current_date - interval '4 months', 'm'),
('ENBL', 'oil commodity stock', current_date - interval '4 months', 'm'),
('ENLC', 'oil commodity stock', current_date - interval '4 months', 'm'),
('EOG', 'oil commodity stock', current_date - interval '4 months', 'm'),
('EPD', 'oil commodity stock', current_date - interval '4 months', 'm'),
('EQT', 'oil commodity stock', current_date - interval '4 months', 'm'),
('ET', 'oil commodity stock', current_date - interval '4 months', 'm'),
('ETRN', 'oil commodity stock', current_date - interval '4 months', 'm'),
('FANG', 'oil commodity stock', current_date - interval '4 months', 'm'),
('FTI', 'oil commodity stock', current_date - interval '4 months', 'm'),
('GPRK', 'oil commodity stock', current_date - interval '4 months', 'm'),
('HAL', 'oil commodity stock', current_date - interval '4 months', 'm'),
('HES', 'oil commodity stock', current_date - interval '4 months', 'm'),
('HP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('KMI', 'oil commodity stock', current_date - interval '4 months', 'm'),
('LIN', 'oil commodity stock', current_date - interval '4 months', 'm'),
('LNG', 'oil commodity stock', current_date - interval '4 months', 'm'),
('MGY', 'oil commodity stock', current_date - interval '4 months', 'm'),
('MPC', 'oil commodity stock', current_date - interval '4 months', 'm'),
('MPLX', 'oil commodity stock', current_date - interval '4 months', 'm'),
('MRO', 'oil commodity stock', current_date - interval '4 months', 'm'),
('MTDR', 'oil commodity stock', current_date - interval '4 months', 'm'),
('MUR', 'oil commodity stock', current_date - interval '4 months', 'm'),
('NBR', 'oil commodity stock', current_date - interval '4 months', 'm'),
('NFE', 'oil commodity stock', current_date - interval '4 months', 'm'),
('NOV', 'oil commodity stock', current_date - interval '4 months', 'm'),
('OAS', 'oil commodity stock', current_date - interval '4 months', 'm'),
('OXY', 'oil commodity stock', current_date - interval '4 months', 'm'),
('PAA', 'oil commodity stock', current_date - interval '4 months', 'm'),
('AM', 'oil commodity stock', current_date - interval '4 months', 'm'),
('APA', 'oil commodity stock', current_date - interval '4 months', 'm'),
('BKR', 'oil commodity stock', current_date - interval '4 months', 'm'),
('BP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('CLB', 'oil commodity stock', current_date - interval '4 months', 'm'),
('CNX', 'oil commodity stock', current_date - interval '4 months', 'm'),
('COP', 'oil commodity stock', current_date - interval '4 months', 'm'),
('CPE', 'oil commodity stock', current_date - interval '4 months', 'm'),
('CQP', 'oil commodity stock', current_date - interval '4 months', 'm')
on conflict (tickersymbol) do nothing
;

-- Buffet liked these at some point or held them
    insert into tickers (tickersymbol, type, lastupdated, importance) values
('NU', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('AMZN', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('STNE', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('DHI', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('LEN', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('AAPL', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('FND', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('DVA', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('NVR', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('MCO', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('SNOW', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('COF', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('ALLY', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('ITOCF', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('V', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('TMUS', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('MA', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('BATRK', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('AXP', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('HPQ', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('CHTR', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('AON', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('LPX', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('JEF', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('VRSN', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('MKL', 'Buffet need-label stock', current_date - interval '4 months', 'm'), -- I like this, like mini berkshire
('GL', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('KR', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('KO', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('PARA', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('SIRI', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('LILA', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('KHC', 'Buffet need-label stock', current_date - interval '4 months', 'm'),
('LILAK', 'Buffet need-label stock', current_date - interval '4 months', 'm')
on conflict (tickersymbol) do nothing
;

-- financials banks
    insert into tickers (tickersymbol, type, lastupdated, importance) values
('C', 'Buffet financial stock', current_date - interval '4 months', 'm'),
('ABCB', 'financial stock', current_date - interval '4 months', 'm'),
('ACGL', 'financial stock', current_date - interval '4 months', 'm'),
('AGNC', 'financial stock', current_date - interval '4 months', 'm'),
('AMTD', 'financial stock', current_date - interval '4 months', 'm'),
('ANAT', 'financial stock', current_date - interval '4 months', 'm'),
('AUB', 'financial stock', current_date - interval '4 months', 'm'),
('BANR', 'financial stock', current_date - interval '4 months', 'm'),
('BGCP', 'financial stock', current_date - interval '4 months', 'm'),
('BHF', 'financial stock', current_date - interval '4 months', 'm'),
('BOKF', 'financial stock', current_date - interval '4 months', 'm'),
('BPOP', 'financial stock', current_date - interval '4 months', 'm'),
('CACC', 'financial stock', current_date - interval '4 months', 'm'),
('CATY', 'financial stock', current_date - interval '4 months', 'm'),
('CBSH', 'financial stock', current_date - interval '4 months', 'm'),
('CFFN', 'financial stock', current_date - interval '4 months', 'm'),
('CIGI', 'financial stock', current_date - interval '4 months', 'm'),
('CINF', 'financial stock', current_date - interval '4 months', 'm'),
('CME', 'financial stock', current_date - interval '4 months', 'm'),
('COLB', 'financial stock', current_date - interval '4 months', 'm'),
('CONE', 'financial stock', current_date - interval '4 months', 'm'),
('CSFL', 'financial stock', current_date - interval '4 months', 'm'),
('CTRE', 'financial stock', current_date - interval '4 months', 'm'),
('CVBF', 'financial stock', current_date - interval '4 months', 'm'),
('EQIX', 'financial stock', current_date - interval '4 months', 'm'),
('ERIE', 'financial stock', current_date - interval '4 months', 'm'),
('ESGR', 'financial stock', current_date - interval '4 months', 'm'),
('ETFC', 'financial stock', current_date - interval '4 months', 'm'),
('EWBC', 'financial stock', current_date - interval '4 months', 'm'),
('FCFS', 'financial stock', current_date - interval '4 months', 'm'),
('FCNCA', 'financial stock', current_date - interval '4 months', 'm'),
('FFBC', 'financial stock', current_date - interval '4 months', 'm'),
('FFIN', 'financial stock', current_date - interval '4 months', 'm'),
('FHB', 'financial stock', current_date - interval '4 months', 'm'),
('FITB', 'financial stock', current_date - interval '4 months', 'm'),
('FMBI', 'financial stock', current_date - interval '4 months', 'm'),
('FRME', 'financial stock', current_date - interval '4 months', 'm'),
('FSV', 'financial stock', current_date - interval '4 months', 'm'),
('FULT', 'financial stock', current_date - interval '4 months', 'm'),
('GBCI', 'financial stock', current_date - interval '4 months', 'm'),
('GLIBA', 'financial stock', current_date - interval '4 months', 'm'),
('GLPI', 'financial stock', current_date - interval '4 months', 'm'),
('HBAN', 'financial stock', current_date - interval '4 months', 'm'),
('HOMB', 'financial stock', current_date - interval '4 months', 'm'),
('HOPE', 'financial stock', current_date - interval '4 months', 'm'),
('HWC', 'financial stock', current_date - interval '4 months', 'm'),
('IBKC', 'financial stock', current_date - interval '4 months', 'm'),
('IBOC', 'financial stock', current_date - interval '4 months', 'm'),
('IBTX', 'financial stock', current_date - interval '4 months', 'm'),
('INDB', 'financial stock', current_date - interval '4 months', 'm'),
('ISBC', 'financial stock', current_date - interval '4 months', 'm'),
('LAMR', 'financial stock', current_date - interval '4 months', 'm'),
('LPLA', 'financial stock', current_date - interval '4 months', 'm'),
('LTXB', 'financial stock', current_date - interval '4 months', 'm'),
('MKTX', 'financial stock', current_date - interval '4 months', 'm'),
('NAVI', 'financial stock', current_date - interval '4 months', 'm'),
('NDAQ', 'financial stock', current_date - interval '4 months', 'm'),
('NGHC', 'financial stock', current_date - interval '4 months', 'm'),
('NMRK', 'financial stock', current_date - interval '4 months', 'm'),
('NTRS', 'financial stock', current_date - interval '4 months', 'm'),
('ONB', 'financial stock', current_date - interval '4 months', 'm'),
('OZK', 'financial stock', current_date - interval '4 months', 'm'),
('PACW', 'financial stock', current_date - interval '4 months', 'm'),
('PBCT', 'financial stock', current_date - interval '4 months', 'm'),
('PCH', 'financial stock', current_date - interval '4 months', 'm'),
('PFG', 'financial stock', current_date - interval '4 months', 'm'),
('PNFP', 'financial stock', current_date - interval '4 months', 'm'),
('PPBI', 'financial stock', current_date - interval '4 months', 'm'),
('REG', 'financial stock', current_date - interval '4 months', 'm'),
('RNST', 'financial stock', current_date - interval '4 months', 'm'),
('ROIC', 'financial stock', current_date - interval '4 months', 'm'),
('SBAC', 'financial stock', current_date - interval '4 months', 'm'),
('SBNY', 'financial stock', current_date - interval '4 months', 'm'),
('SBRA', 'financial stock', current_date - interval '4 months', 'm'),
('SEIC', 'financial stock', current_date - interval '4 months', 'm'),
('SFBS', 'financial stock', current_date - interval '4 months', 'm'),
('SFNC', 'financial stock', current_date - interval '4 months', 'm'),
('SIGI', 'financial stock', current_date - interval '4 months', 'm'),
('SIVB', 'financial stock', current_date - interval '4 months', 'm'),
('SLM', 'financial stock', current_date - interval '4 months', 'm'),
('SSB', 'financial stock', current_date - interval '4 months', 'm'),
('TCBI', 'financial stock', current_date - interval '4 months', 'm'),
('TCF', 'financial stock', current_date - interval '4 months', 'm'),
('TFSL', 'financial stock', current_date - interval '4 months', 'm'),
('TOWN', 'financial stock', current_date - interval '4 months', 'm'),
('TREE', 'financial stock', current_date - interval '4 months', 'm'),
('TRMK', 'financial stock', current_date - interval '4 months', 'm'),
('TROW', 'financial stock', current_date - interval '4 months', 'm'),
('UBSI', 'financial stock', current_date - interval '4 months', 'm'),
('UCBI', 'financial stock', current_date - interval '4 months', 'm'),
('UMBF', 'financial stock', current_date - interval '4 months', 'm'),
('UMPQ', 'financial stock', current_date - interval '4 months', 'm'),
('UNIT', 'financial stock', current_date - interval '4 months', 'm'),
('VLY', 'financial stock', current_date - interval '4 months', 'm'),
('WAFD', 'financial stock', current_date - interval '4 months', 'm'),
('WSBC', 'financial stock', current_date - interval '4 months', 'm'),
('WSFS', 'financial stock', current_date - interval '4 months', 'm'),
('WTFC', 'financial stock', current_date - interval '4 months', 'm'),
('ZG', 'financial stock', current_date - interval '4 months', 'm'),
('ZION', 'financial stock', current_date - interval '4 months', 'm')
on conflict (tickersymbol) do nothing
;

    insert into tickers (tickersymbol, type, lastupdated, importance) values
('TLT', 'bond macro ETF', current_date - interval '4 months', 'm'),
('BRK-B', 'Buffet stock', current_date - interval '4 months', 'm'),
-- Japan
('MUFG', 'Japan financial stock', current_date - interval '4 months', 'm'),
('SMFG', 'Japan financial stock', current_date - interval '4 months', 'm'),
('HTCR', 'Japan tech stock', current_date - interval '4 months', 'm'),
('NMR', 'Japan financial stock', current_date - interval '4 months', 'm'),
('IX', 'Japan financial stock', current_date - interval '4 months', 'm'),
('MFG', 'Japan financial stock', current_date - interval '4 months', 'm'),
('TAK', 'Japan pharma stock', current_date - interval '4 months', 'm'),
('SONY', 'Japan tech industrial stock', current_date - interval '4 months', 'm'),
('TM', 'Japan industrial stock', current_date - interval '4 months', 'm'),
('MSBHF', 'Japan financial Buffet stock', current_date - interval '4 months', 'm'),
('EBCOY', 'Japan industrial stock', current_date - interval '4 months', 'm'),
('MRM', 'Japan industrial stock', current_date - interval '4 months', 'm');

