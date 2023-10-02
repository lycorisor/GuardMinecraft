package _const

// serverip serverport privatekey publickey reserved
const ProxyDefaultConfig = `{
    "log": {
        "level": "error",
        "timestamp": true
    },
    "dns": {
        "servers": [
            {
                "tag": "dns_proxy",
                "address": "udp://1.0.0.1",
                "detour": "direct"
            },
            {
                "tag": "dns_direct",
                "address": "https://dns.alidns.com/dns-query",
                "address_resolver": "dns_resolver",
                "detour": "direct"
            },
            {
                "tag": "dns_resolver",
                "address": "223.5.5.5",
                "detour": "direct"
            }
        ],
        "rules": [
            {
                "domain_suffix": [
                    "mojang.com",
                    "minecraft.net",
                    "minecraftservices.com"
                ],
                "server": "dns_proxy"
            },
            {
                "outbound": [
                    "any"
                ],
                "server": "dns_resolver"
            }
        ]
    },
    "route": {
        "rules": [
            {
                "domain_suffix": [
                    "mojang.com",
                    "minecraft.net",
                    "minecraftservices.com"
                ],
                "outbound": "cfwarp"
            },
            {
                "protocol": "dns",
                "outbound": "dns-out"
            }
        ],
 		"final": "direct",
        "auto_detect_interface": true
    },
	"inbounds": [
        {
            "type": "tun",
            "tag": "tun-in",
            "inet4_address": "10.0.0.6/30",
            "auto_route": true,
            "stack": "system",
            "sniff": true,
            "sniff_override_destination": false
        }
    ],
	"outbounds": [
        {
            "type": "direct",
            "tag": "direct"
        },
        {
            "type": "wireguard",
            "tag": "cfwarp",
            "server": "%s",
            "server_port": %d,
            "local_address": ["172.16.0.2/32"],
            "private_key": "%s",
            "peer_public_key": "%s",
            "reserved": %s,
            "mtu": 1280
	    },
        {
            "type": "dns",
            "tag": "dns-out"
        }
    ]
}`
