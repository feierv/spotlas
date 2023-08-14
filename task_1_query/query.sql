SELECT
    spot_name AS name,
    domain,
    COUNT(*) AS count_for_domain
FROM (
    SELECT
        name AS spot_name,
        regexp_replace(website, '^(https?://)?(www\.)?([^/]+).*$', '\3') AS domain
    FROM "MY_TABLE"
    WHERE website IS NOT NULL
) AS extracted_domains
GROUP BY spot_name, domain
ORDER BY count_for_domain DESC;
