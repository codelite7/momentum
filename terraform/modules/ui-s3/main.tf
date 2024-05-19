locals {
  # content type conversion map, to override s3 content type using file extension
  content_types = {
    "css"  = "text/css"
    "eot"  = "application/vnd.ms-fontobject"
    "html" = "text/html"
    "ico"  = "image/vnd.microsoft.icon"
    "jpg"  = "image/jpeg"
    "js"   = "application/javascript"
    "json" = "application/json"
    "map"  = "application/json"
    "png"  = "image/png"
    "svg"  = "image/svg+xml"
    "ttf"  = "font/ttf"
    "txt"  = "text/plain"
    "webp" = "image/webp"
    "woff" = "font/woff"
  }
}

resource "aws_s3_bucket" "bucket" {
  bucket = var.domain
  force_destroy = true
}

resource "aws_s3_bucket_website_configuration" "site" {
  bucket = aws_s3_bucket.bucket.id
  index_document {
    suffix = "index.html"
  }
  error_document {
    key = "index.html"
  }
}

resource "aws_s3_bucket_public_access_block" "public_access_block" {
  bucket = aws_s3_bucket.bucket.id
  block_public_acls = false
  block_public_policy = false
#   ignore_public_acls = false
#   restrict_public_buckets = false
}

resource "aws_s3_object" "upload_object" {
  for_each = fileset(var.content-path, "**")
  bucket = aws_s3_bucket.bucket.id
  key = each.value
  source = "${var.content-path}/${each.value}"
  content_type = lookup(local.content_types, element(split(".", each.value), length(split(".", each.value)) - 1), "text/plain")
  etag = filemd5("${var.content-path}/${each.value}")
}

resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.bucket.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect    = "Allow"
        Principal = "*"
        Action = [
          "s3:GetObject"
        ]
        Resource = [
          "${aws_s3_bucket.bucket.arn}/*"
        ]
      }
    ]
  })
}

resource "cloudflare_record" "ui-cname" {
  zone_id = var.cloudflareZoneId
  type = "CNAME"
  name = var.domain
  value = aws_s3_bucket_website_configuration.site.website_endpoint
  proxied = var.proxied
  ttl = var.ttl
}