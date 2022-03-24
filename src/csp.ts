import { Client } from "@elastic/elasticsearch";
import { IncomingMessage, ServerResponse } from "micri";

const client = new Client({ node: 'http://localhost:9200' });

export default async function csp(_req: IncomingMessage, res: ServerResponse) {
  await client.index({
    index: 'game-of-thrones',
    document: {
      character: 'Ned Stark',
      quote: 'Winter is coming.'
    }
  });
  res.writeHead(413, {
    "X-Backoff": "Time to slow down"
  });
  res.end();
}