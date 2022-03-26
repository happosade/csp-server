import { IncomingMessage, ServerResponse } from "micri";

export default async function wrkHello(_req: IncomingMessage, res: ServerResponse) {
  res.writeHead(421, {
    'X-custom': "X-HEADER-D"
  });
  res.write("This is not the page you're looking for...\n");
  res.end();
}