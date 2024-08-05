const { log } = require("console");
const fs = require("fs")

if (process.argv.length < 3) {
  console.log("too few args")
  return
}

const openapi = JSON.parse(fs.readFileSync(process.argv[2], 'utf8'));

openapi["openapi"] = "3.0.3"
// delete openapi["paths"]["/api/auth/login"]["post"]
delete openapi["components"]["schemas"]["SmppSmsProviderConfiguration"]["allOf"][1]["properties"]["codingScheme"]

fs.writeFileSync(process.argv[2], JSON.stringify(openapi), 'utf8')
