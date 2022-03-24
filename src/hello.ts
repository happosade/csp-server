import { IncomingMessage, ServerResponse } from "micri";

export default async function wrkHello(_req: IncomingMessage, res: ServerResponse) {
  res.writeHead(200, {
    'X-custom': "X-HEADER-D"
  });
  res.write("Hello! world!!!\n");
  res.end();
}