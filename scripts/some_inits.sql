
-- stocks I like
insert into tickers (tickersymbol, lastupdated, importance) values
('IBM', current_date - interval '4 months', 'm'), -- the demos use it
    ('CLF', current_date - interval '4 months', 'm'),
    ('MT', current_date - interval '4 months', 'm'),
    ('X', current_date - interval '4 months', 'm'),
    ('STLD', current_date - interval '4 months', 'm'),
    ('GGB', current_date - interval '4 months', 'm'),
    ('TX', current_date - interval '4 months', 'm'),
    ('RS', current_date - interval '4 months', 'm'),
    ('EWZ', current_date - interval '4 months', 'm'),
    ('PBR-A', current_date - interval '4 months', 'm'),
    ('MRO', current_date - interval '4 months', 'm'),
    ('OXY', current_date - interval '4 months', 'm');

insert into datasources (sourceid, sourcename, sourceurl) values (1, 'Alphavantage', 'https://www.alphavantage.co');
