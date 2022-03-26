export interface CspReport {
  age?: number;
  body: Body;
  type?: string;
  url?: string;
  user_agent?: string;
}
export interface Body {
  blockedURL?: string;
  disposition?: string;
  documentURL?: string;
  effectiveDirective?: string;
  lineNumber?: number;
  originalPolicy?: string;
  referrer?: string;
  sourceFile?: string;
  statusCode?: number;
  [key: string]: unknown;
}
