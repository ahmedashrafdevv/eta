GO
ALTER PROCEDURE [dbo].[StkTr01ConvertInvoice](@Serial int )
	AS
    BEGIN
        DECLARE @internalID VARCHAR(100)
        DECLARE @storeCode int
        DECLARE @dateTimeIssued datetime
        DECLARE @totalDiscountAmount REAL
        DECLARE @totalAmount REAL
        DECLARE @totalTax REAL
        DECLARE @accountSerial int
		
        SELECT  @internalID =   DocNo ,
                @storeCode = StoreCode , 
                @dateTimeIssued = DocDate , 
                @totalDiscountAmount = Discount , 
                @totalAmount = TotalCash , 
                @totalTax = ISNULL(TotalTax , 0)  , 
                @accountSerial = AccountSerial FROM StkTr01 WHERE Serial = @Serial
        INSERT INTO StkTrEInvoiceHead (internalID ,storeCode ,dateTimeIssued ,totalDiscountAmount ,totalAmount ,totalTax ,stkTr01Serial ,accountSerial) VALUES (@internalID ,@storeCode ,@dateTimeIssued ,@totalDiscountAmount ,@totalAmount ,@totalTax ,@Serial ,@accountSerial)

        SET @headSerial = SCOPE_IDENTITY();

        INSERT INTO StkTrEInvoiceDetails ( headeSerial , itemCode  , unitValue , quantity , totalTaxableFees , itemsDiscount , itemSerial )
        SELECT @headSerial , ISNULL(BarCodeUsed, '123') , Price , Qnt , Tax , ISNULL(Discount,0 ), ItemSerial 
        FROM    StkTr02 WHERE HeadSerial = @Serial  


		UPDATE StkTr01 SET EtaConverted = 1 WHERE "Serial" = @Serial
		SELECT SCOPE_IDENTITY() id
    END


GO
CREATE PROCEDURE StkTrEInvoiceHeadList(@posted BIT = NULL)
	AS
    BEGIN
        SELECT  h.Serial ,  h.internalID ,  h.storeCode , totalDiscountAmount , totalAmount , h.totalTax ,  h.stkTr01Serial  FROM StkTrEInvoiceHead h WHERE deletedAt IS NULL AND h.posted = CASE WHEN @posted IS NULL THEN h.posted ELSE @posted END; 
    END


GO
ALTER PROCEDURE StkTrEInvoiceFind(@serial INT)
	AS
    BEGIN
        DECLARE @totalSalesAmount FLOAT
        DECLARE @totalItemsDiscountAmount FLOAT
        DECLARE @issuerRegistrationId NVARCHAR(100)
        DECLARE @activityCode NVARCHAR(100)
        DECLARE @issuerType NVARCHAR(1)
        DECLARE @issuerName NVARCHAR(100)

        SELECT @issuerRegistrationId = EtaRegistrationId ,@activityCode = EtaActivityCode , @issuerType = EtaType , @issuerName = ComName FROM ComInfo
        SELECT @totalSalesAmount = SUM(quantity * unitValue) , @totalItemsDiscountAmount = SUM(itemsDiscount) FROM StkTrEInvoiceDetails WHERE headSerial = @serial
        SELECT  
            h.dateTimeIssued , h.internalID ,@totalSalesAmount totalSalesAmount ,@totalItemsDiscountAmount totalItemsDiscountAmount,
            h.totalDiscountAmount , (h.totalDiscountAmount - @totalItemsDiscountAmount) ExtraDiscountAmount ,
            (@totalSalesAmount - h.totalDiscountAmount) netAmount  ,h.totalAmount ,
            h.accountSerial ,  h.storeCode ,
            @issuerRegistrationId,@issuerName ,@issuerType , @activityCode
        FROM StkTrEInvoiceHead h
        WHERE h.Serial = @serial 
    END

GO
ALTER PROCEDURE StkTrEInvoiceFindItems(@serial INT)
	AS
    BEGIN

    SELECT itemType , itemCode , unitType , quantity , unitValue ,totalTaxableFees ,itemsDiscount, (quantity * unitValue) salesTotal , ((quantity * unitValue) + totalTaxableFees) total  , ((quantity * unitValue) + totalTaxableFees - itemsDiscount) netTotal  FROM StkTrEInvoiceDetails WHERE headSerial = @serial
END
GO
ALTER PROCEDURE StkTrEInvoiceFindRecieverAddress(@accountSerial INT)
	AS
    BEGIN
        SELECT ad.* , acc.AccountName , acc.EtaType , acc.EtaId FROM  EtaAddresses ad  JOIN AccMs01 acc ON acc.AddressSerial = ad.Serial WHERE acc.Serial = @accountSerial
    END
GO
ALTER PROCEDURE StkTrEInvoiceFindIssuerAddress(@storeCode INT)
	AS
    BEGIN
        SELECT ad.* FROM  EtaAddresses ad JOIN StoreCode st ON st.AddressSerial = ad.Serial WHERE st.StoreCode = @storeCode
    END




GO
ALTER PROCEDURE [dbo].[StkTr01ListByTransSerialAndConverted] (@TransSerial INT , @Converted SMALLINT = -1 , @StoreCode INT = 0 )
	AS
    BEGIN
        SELECT  o.Serial ,  o.DocNo ,  o.DocDate ,  ISNULL(o.Discount , 0) Discount  ,  o.TotalCash ,
			  ISNULL(o.TotalTax , 0) TotalTax    
	    FROM StkTr01 o 
		WHERE o.TotalCash IS NOT NULL 
		AND ISNULL(EtaConverted , 0) = CASE WHEN @Converted = -1 THEN ISNULL(EtaConverted , 0) ELSE  @Converted END 
		AND TransSerial = @TransSerial
		AND StoreCode = CASE WHEN @StoreCode = 0 THEN StoreCode ELSE @StoreCode END
    END






