-- Create "api_tokens" table
CREATE TABLE "api_tokens" ("id" uuid NOT NULL, "description" character varying NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "expires_at" timestamptz NULL, "revoked_at" timestamptz NULL, "organization_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "api_tokens_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
