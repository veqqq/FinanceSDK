insert into tickers (tickersymbol, lastupdated, importance, type) values
    -- these represent the majority of volume
    ('FOREX EUR JPY', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX USD GBP', current_date - interval '4 months', 'm', 'Forex'), -- pound dollar
    ('FOREX USD CHF', current_date - interval '4 months', 'm', 'Forex'), -- dollar swissy
    ('FOREX USD CAD', current_date - interval '4 months', 'm', 'Forex'), -- dollar loonie
    ('FOREX AUD USD', current_date - interval '4 months', 'm', 'Forex'), -- Aussie dollar
    ('FOREX NZD USD', current_date - interval '4 months', 'm', 'Forex'), -- kiwi dollar
    ('FOREX USD EUR', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX CAD JPY', current_date - interval '4 months', 'm', 'Forex'), -- correlates to oil prices
    ('FOREX USD JPY', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX EUR CAD', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX EUR AUD', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX EUR CHF', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX AUD CAD', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX GBP CHF', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX GBP JPY', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX CHF JPY', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX AUD JPY', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX AUD NZD', current_date - interval '4 months', 'm', 'Forex'),
    -- I think these are important/useful
    ('FOREX USD CNY', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX USD HKD', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX USD INR', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX BRL USD', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX USD PEN', current_date - interval '4 months', 'm', 'Forex'), -- peru
    ('FOREX RUB TRY', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX RUB EUR', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX USD PLN', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX USD KRW', current_date - interval '4 months', 'm', 'Forex'),
    ('FOREX USD THB', current_date - interval '4 months', 'm', 'Forex')
;