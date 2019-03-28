INSERT INTO `urls` (`id`, `created_at`, `updated_at`, `deleted_at`, `url`, `state`, `status`, `download_http_code`) VALUES (1, '2019-03-26 22:28:11', '2019-03-27 20:59:32', NULL, 'https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html', 1, 1, 0);
INSERT INTO `urls` (`id`, `created_at`, `updated_at`, `deleted_at`, `url`, `state`, `status`, `download_http_code`) VALUES (2, '2019-03-26 22:28:11', '2019-03-27 20:59:19', NULL, 'https://vietnamnet.vn/vn/cong-nghe/ung-dung/cach-su-dung-google-maps-de-giam-sat-vi-tri-cua-tre-nho-514378.html', 1, 1, 0);
INSERT INTO `urls` (`id`, `created_at`, `updated_at`, `deleted_at`, `url`, `state`, `status`, `download_http_code`) VALUES (3, '2019-03-26 22:28:11', '2019-03-27 20:59:35', NULL, 'https://vietnamnet.vn/vn/doi-song/song-la/ba-pham-thi-yen-xin-loi-gia-dinh-nu-sinh-giao-ga-o-dien-bien-516657.html', 1, 1, 0);
INSERT INTO `urls` (`id`, `created_at`, `updated_at`, `deleted_at`, `url`, `state`, `status`, `download_http_code`) VALUES (4, '2019-03-26 22:28:11', '2019-03-27 20:59:22', NULL, 'https://www.thesaigontimes.vn/274112/tiep-tuc-da-cat-giam-thu-tuc-hanh-chinh.html', 1, 1, 0);
INSERT INTO `urls` (`id`, `created_at`, `updated_at`, `deleted_at`, `url`, `state`, `status`, `download_http_code`) VALUES (5, '2019-03-26 22:28:11', '2019-03-27 20:59:38', NULL, 'https://www.thesaigontimes.vn/274105/phe-lieu-va-logistics.html', 1, 1, 0);


-- Reset de crawl lai
update urls
set urls.state = 1, urls.`status` = 1;

-- Reset db vi url_id la unique
delete from articles;

-- Minh hoa ve cach chay tren nhieu instance
select *
from urls
where id % 3 = 0

select *
from urls
where id % 3 = 1

select *
from urls
where id % 3 = 2