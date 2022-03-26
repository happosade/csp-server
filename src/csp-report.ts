import { Client } from "@elastic/elasticsearch";
import { IncomingMessage, ServerResponse } from "micri";

const client = new Client({
  node: process.env.ES_NODE || 'http://localhost:9200'
});

export default async function csp(_req: IncomingMessage, res: ServerResponse) {
  await client.index({
    index: 'game-of-thrones',
    document: {
      character: 'Ned Stark',
      quote: 'Winter is coming.'
    }
  });
  res.writeHead(200);
  res.write(_req.headers["user-agent"] + "\n"); // Figuring out how this works
  res.end();
}