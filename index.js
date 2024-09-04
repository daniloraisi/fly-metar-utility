const http = require("http");
const url = require("url");
const puppeteer = require("puppeteer");

http
  .createServer(async function (req, res) {
    const queryObject = url.parse(req.url, true).query;

    if (!queryObject.icao) {
      res.end("Informe um ICAO");
      return;
    }

    const metarUrl = `https://metar-taf.com/${queryObject.icao}`;
    const browser = await puppeteer.launch({
      args: ["--no-sandbox", "disable-setupid-sandbox"],
    });
    const page = await browser.newPage();

    try {
      await page.goto(metarUrl);
    } catch (e) {
      console.log(e);
      res.end();
      return;
    }

    const expression =
      "//code[@class='text-white d-block']";
    const elements = await page.$x(expression);

    try {
      await page.waitForXPath(expression, { timeout: 3000 });
    } catch (error) {
      console.log(error);
      res.end();
      return;
    }

    const metar = await page.evaluate((el) => el.innerHTML, elements[0]);
    console.log(metar);

    await page.close();
    await browser.close();

    res.end(metar);
  })
  .listen(process.env.PORT || 3000, process.env.HOST || "localhost");
