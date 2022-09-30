<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/products.proto

namespace GPBMetadata\Proto\MemberExtension;

class Products
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
%proto/member-extension/products.protomemberExtension"�
Product

id (	
category (	
merchant_id (	
product_name (	
product_code (	
images (	
videos (	
graphic_detail (	
describe	 (	

sale_begin
 (
sale_end (+
carriage (2.memberExtension.Carriage
status
extra (	
sku (	
category_id (	
weight_value (
point (
price (

line_price (
weight (
quantity (
surplus_quantity (
distribute_coupon_id (	

benefit_id (	
benefit_valid_day (
benefit_level_limit (	%
distribute_coupon_category_id (	
coupon_category (	
benefit_division_limit (	 
benefit_level_limit_name (	#
benefit_division_limit_name  (	"\'
Carriage
cash (
point ("A
CreateProductRequest)
product (2.memberExtension.Product"P
CreateProductResponse

error_code (

data (	"�
GetProductsRequest
name (	
category (	
status (	
point_begin (
	point_end (
category_id (	
offset (
limit (
order_field	 (	
	order_asc
 (
on_sale ("p
GetProductsResponse

error_code (

data (2 .memberExtension.GetProductsData"L
GetProductsData*
products (2.memberExtension.Product
total ("7
UpdateProductQuantityRequest

id (	
num ("A
UpdateProductRequest)
product (2.memberExtension.Product",
GetProductQuantityBillsRequest

id (	"�
GetProductQuantityBillsResponse

error_code (

data (2,.memberExtension.GetProductQuantityBillsData"a
GetProductQuantityBillsData
code (	

created_at (
change (
quantity (" 
ShowProductRequest

id (	"h
ShowProductResponse

error_code (

data (2.memberExtension.Product"8
UpdateProductStatusRequest

id (	
status (	B+Z
        , true);

        static::$is_initialized = true;
    }
}
