# 简介

本项目每周四自动生成 GeoIP 文件

### V2Ray dat 格式路由规则文件

- **geoip.dat**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip.dat](https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip.dat)
- **geoip.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip.dat.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip.dat.sha256sum)
- **geoip-only-cn-private.dat**（精简版 GeoIP，只包含 `geoip:cn` 和 `geoip:private`）：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-only-cn-private.dat](https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-only-cn-private.dat)
- **geoip-only-cn-private.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-only-cn-private.dat.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-only-cn-private.dat.sha256sum)
- **geoip-asn.dat**（精简版 GeoIP，只包含上述新增类别）：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-asn.dat](https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-asn.dat)
- **geoip-asn.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-asn.dat.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/geoip-asn.dat.sha256sum)
- **cn.dat**（精简版 GeoIP，只包含 `geoip:cn`）：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/cn.dat](https://raw.githubusercontent.com/Tiiwoo/geoip/release/cn.dat)
- **cn.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/cn.dat.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/cn.dat.sha256sum)
- **private.dat**（精简版 GeoIP，只包含 `geoip:private`）：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/private.dat](https://raw.githubusercontent.com/Tiiwoo/geoip/release/private.dat)
- **private.dat.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/private.dat.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/private.dat.sha256sum)

### MaxMind mmdb 格式文件

- **Country.mmdb**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country.mmdb](https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country.mmdb)
- **Country.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country.mmdb.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country.mmdb.sha256sum)
- **Country-only-cn-private.mmdb**（精简版 GeoIP，只包含 `GEOIP,CN` 和 `GEOIP,PRIVATE`）：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-only-cn-private.mmdb](https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-only-cn-private.mmdb)
- **Country-only-cn-private.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-only-cn-private.mmdb.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-only-cn-private.mmdb.sha256sum)
- **Country-asn.mmdb**（精简版 GeoIP，只包含上述新增类别）：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-asn.mmdb](https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-asn.mmdb)
- **Country-asn.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-asn.mmdb.sha256sum](https://raw.githubusercontent.com/Tiiwoo/geoip/release/Country-asn.mmdb.sha256sum)

### 支持的格式

关于每种格式所支持的配置选项，查看本项目 [`config-example.json`](https://github.com/Tiiwoo/geoip/blob/HEAD/config-example.json) 文件。

支持的 `input` 输入格式：

- **text**：纯文本 IP 和 CIDR（例如：`1.1.1.1` 或 `1.0.0.0/24`）
- **private**：局域网和私有网络 CIDR（例如：`192.168.0.0/16` 和 `127.0.0.0/8`）
- **cutter**：用于裁剪前置步骤中的数据
- **v2rayGeoIPDat**：V2Ray GeoIP dat 格式（`geoip.dat`）
- **maxmindMMDB**：MaxMind mmdb 数据格式（`GeoLite2-Country.mmdb`）
- **maxmindGeoLite2CountryCSV**：MaxMind GeoLite2 country CSV 数据（`GeoLite2-Country-CSV.zip`）
- **clashRuleSetClassical**：[classical 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#classical)
- **clashRuleSet**：[ipcidr 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#ipcidr)
- **surgeRuleSet**：[Surge RuleSet](https://manual.nssurge.com/rule/ruleset.html)

支持的 `output` 输出格式：

- **text**：纯文本 CIDR（例如：`1.0.0.0/24`）
- **v2rayGeoIPDat**：V2Ray GeoIP dat 格式（`geoip.dat`，适用于 [V2Ray](https://github.com/v2fly/v2ray-core)、[Xray-core](https://github.com/XTLS/Xray-core) 和 [Trojan-Go](https://github.com/p4gefau1t/trojan-go)）
- **maxmindMMDB**：MaxMind mmdb 数据格式（`GeoLite2-Country.mmdb`，适用于 [Clash](https://github.com/Dreamacro/clash) 和 [Leaf](https://github.com/eycorsican/leaf)）
- **clashRuleSetClassical**：[classical 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#classical)
- **clashRuleSet**：[ipcidr 类型的 Clash RuleSet](https://github.com/Dreamacro/clash/wiki/premium-core-features#ipcidr)
- **surgeRuleSet**：[Surge RuleSet](https://manual.nssurge.com/rule/ruleset.html)

### 注意事项

由于 MaxMind mmdb 文件格式的限制，当不同列表的 IP 或 CIDR 数据有交集或重复项时，后写入的列表的 IP 或 CIDR 数据会覆盖（overwrite）之前已写入的列表的数据。譬如，IP `1.1.1.1` 同属于列表 `AU` 和列表 `Cloudflare`。如果 `Cloudflare` 在 `AU` 之后写入，则 IP `1.1.1.1` 归属于列表 `Cloudflare`。

为了确保某些指定的列表、被修改的列表一定囊括属于它的所有 IP 或 CIDR 数据，可在 `output` 输出格式为 `maxmindMMDB` 的配置中增加选项 `overwriteList`，该选项中指定的列表会在最后逐一写入，列表中最后一项优先级最高。若已设置选项 `wantedList`，则无需设置 `overwriteList`。`wantedList` 中指定的列表会在最后逐一写入，列表中最后一项优先级最高。
