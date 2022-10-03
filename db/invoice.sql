USE MCSTREE




GO
ALTER PROCEDURE StkTr01ListByTransSerialStoreConvertedDate (@transSerial INT , @converted SMALLINT = NULL , @storeCode INT = 0 , @fromDate VARCHAR(20) = '', @toDate VARCHAR(20) = '')
	AS
    BEGIN
        SELECT  o.Serial ,  o.DocNo ,  o.DocDate ,  ISNULL(o.Discount , 0) Discount  ,  o.TotalCash ,
			  ISNULL(o.TotalTax , 0) TotalTax  , ISNULL(o.EtaConverted , 0) EtaConverted
	    FROM StkTr01 o 
		WHERE o.TotalCash IS NOT NULL 
		AND ISNULL(EtaConverted , 0) = CASE WHEN @Converted IS NULL THEN ISNULL(EtaConverted , 0) ELSE  @Converted END 
		AND TransSerial = @TransSerial
		AND StoreCode = CASE WHEN @StoreCode = 0 THEN StoreCode ELSE @StoreCode END
        AND o.DocDate >= CASE WHEN @fromDate = '' THEN o.DocDate ELSE CONVERT(DATETIME, @fromDate, 102) END
		AND o.DocDate <= CASE WHEN @toDate = '' THEN o.DocDate ELSE CONVERT(DATETIME, @toDate, 102) END
    END

GO
ALTER PROCEDURE StkTr01ConvertInvoice (@Serial INT)
	AS
    BEGIN
        DECLARE @internalID INT
		DECLARE @transSerial INT
        DECLARE @storeCode int
        DECLARE @dateTimeIssued datetime
        DECLARE @totalDiscountAmount REAL
        DECLARE @totalAmount REAL
        DECLARE @totalTax REAL
        DECLARE @accountSerial int
		DECLARE @headSerial int
		


		
        SELECT  
                @storeCode = StoreCode , 
				@transSerial = TransSerial,
                @dateTimeIssued = DocDate , 
                @totalDiscountAmount = Discount , 
                @totalAmount = TotalCash , 
                @totalTax = ISNULL(TotalTax , 0)  , 
                @accountSerial = AccountSerial FROM StkTr01 WHERE Serial = @Serial

		SET @internalID = (SELECT ISNULL(MAX(internalID) , 1) FROM StkTrEInvoiceHead WHERE storeCode = @storeCode AND TransSerial = @transSerial)
        INSERT INTO StkTrEInvoiceHead (internalID , TransSerial ,storeCode ,dateTimeIssued ,totalDiscountAmount ,totalAmount ,totalTax ,stkTr01Serial ,accountSerial)
        VALUES (@internalID ,@transSerial ,@storeCode ,@dateTimeIssued ,@totalDiscountAmount ,@totalAmount ,@totalTax ,@Serial ,@accountSerial)

        SET @headSerial = SCOPE_IDENTITY();

        INSERT INTO StkTrEInvoiceDetails ( headSerial , itemCode  , unitValue , quantity , totalTaxableFees , itemsDiscount , itemSerial )
        SELECT @headSerial , ISNULL(BarCodeUsed, '123') , Price , Qnt , Tax , ISNULL(Discount,0 ), ItemSerial 
        FROM    StkTr02 WHERE HeadSerial = @Serial  


		UPDATE StkTr01 SET EtaConverted = 1 WHERE "Serial" = @Serial
		SELECT SCOPE_IDENTITY() id
    END


GO
ALTER PROCEDURE StkTr03ConvertInvoice (@Serial INT)
	AS
    BEGIN
        DECLARE @internalID INT
		DECLARE @transSerial INT
        DECLARE @storeCode int
        DECLARE @dateTimeIssued datetime
        DECLARE @totalDiscountAmount REAL
        DECLARE @totalAmount REAL
        DECLARE @totalTax REAL
        DECLARE @accountSerial int
		DECLARE @headSerial int
				
        SELECT  
                @storeCode = StoreCode , 
                @dateTimeIssued = DocDate , 
                @totalDiscountAmount = Discount , 
                @totalAmount = TotalCash , 
                @totalTax = ISNULL(SaleTax , 0)  , 
                @accountSerial = AccountSerial FROM StkTr03 WHERE Serial = @Serial

		SET @internalID = (SELECT ISNULL(MAX(internalID) , 1) FROM StkTrEInvoiceHead WHERE storeCode = @storeCode AND TransSerial = @transSerial)
        INSERT INTO StkTrEInvoiceHead (internalID , TransSerial ,storeCode ,dateTimeIssued ,totalDiscountAmount ,totalAmount ,totalTax ,stkTr01Serial ,accountSerial)
        VALUES (@internalID ,25 ,@storeCode ,@dateTimeIssued ,@totalDiscountAmount ,@totalAmount ,@totalTax ,@Serial ,@accountSerial)

        SET @headSerial = SCOPE_IDENTITY();

        INSERT INTO StkTrEInvoiceDetails ( headSerial , itemCode  , unitValue , quantity , totalTaxableFees , itemsDiscount , itemSerial , storeCode )
        SELECT @headSerial , ISNULL(BarCodeUsed, '123') , Price , Qnt , Tax , ISNULL(Discount,0 ), ItemSerial , @storeCode
        FROM    StkTr04 o WHERE HeadSerial = @Serial  

		
		SELECT SCOPE_IDENTITY() id
    END

