
-- stocks I like
insert into tickers (tickersymbol, lastupdated, importance) values
('IBM', current_date - interval '4 months', 'm'), -- demos use it
    ('CLF', current_date - interval '4 months', 'm'),
    ('MT', current_date - interval '4 months', 'm'),
    ('X', current_date - interval '4 months', 'm'),
    ('STLD', current_date - interval '4 months', 'm'),
    ('GGB', current_date - interval '4 months', 'm'),
    ('TX', current_date - interval '4 months', 'm'),
    ('RS', current_date - interval '4 months', 'm'),
    -- energy
    ('PBR-A', current_date - interval '4 months', 'm'),
    ('MRO', current_date - interval '4 months', 'm'),
    ('OXY', current_date - interval '4 months', 'm'),
    ('LEU', current_date - interval '4 months', 'm'),
    -- banks
    ('WAL', current_date - interval '4 months', 'm'),
    ('BDORY', current_date - interval '4 months', 'm')
;

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
    ('NGE', current_date - interval '4 months', 'q', 'Country ETF'), -- Nigeria
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