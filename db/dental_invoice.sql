GO

CREATE TABLE [dbo].[StkTrInvoiceHead](
	[Serial] [int] IDENTITY(1,1) NOT NULL,
	[DocNo] [int] NULL,
	[DocDate] [datetime] NULL,
	[TransSerial] [int] NULL,
	[StoreCode] [smallint] NULL,
	[AccountSerial] [int] NULL,
	[TotalCash] [real] NULL,
	[OrderNo] [int] NULL,
	[IsPrinted] [bit] NULL,
	[IsPosted] [bit] NULL,
	[StkTr01Serial] [int] NULL,
	[delivery_Date] [datetime] NULL,
	[BuyOrSale] [bit] NULL,
	[ShipCode] [int] NULL,
	[CurCode] [int] NULL,
	[CurRatio] [money] NULL,
	[TotalByCur] [money] NULL,
	[Vat] [money] NULL,
	[EmpCode] [int] NULL,
	[AuditCode] [int] NULL,
	[Reserved] [bit] NOT NULL,
	[DeliveryFee] [real] NOT NULL,
	[DriverName] [varchar](100) NULL,
	[CloseDate] [datetime] NULL,
	[DeletedAt] [datetime] NULL,
	[Deleted] [bit] NULL
) ON [PRIMARY]

GO

SET ANSI_PADDING OFF
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT (getdate()) FOR [DocDate]
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT ((0)) FOR [IsPrinted]
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT ((0)) FOR [IsPosted]
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT ((0)) FOR [Vat]
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT ((1)) FOR [Reserved]
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT ((0)) FOR [DeliveryFee]
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT (NULL) FOR [DeletedAt]
GO

ALTER TABLE [dbo].[StkTrInvoiceHead] ADD  DEFAULT ((0)) FOR [Deleted]
GO


GO

CREATE TABLE [dbo].[StkTrInvoiceDetails](
	[Serial] [int] IDENTITY(1,1) NOT NULL,
	[HeadSerial] [int] NOT NULL,
	[ItemSerial] [int] NOT NULL,
	[ItemName] [nvarchar](200) NOT NULL,
	[Qnt] [real] NULL,
	[Price] [real] NULL,
	[MinorPerMajor] [int] NULL,
	[CurCode] [int] NULL,
	[CurRatio] [money] NULL,
	[PriceByCur] [money] NULL,
	[QntAntherUnit] [real] NULL,
	[PriceMax] [money] NULL,
	[PriceMin] [money] NULL,
	[Branch] [real] NULL
) ON [PRIMARY]

GO

ALTER TABLE [dbo].[StkTrInvoiceDetails] ADD  CONSTRAINT [DF__StkTrInvo__CurCo__473666D9]  DEFAULT ((0)) FOR [CurCode]
GO

ALTER TABLE [dbo].[StkTrInvoiceDetails] ADD  CONSTRAINT [DF__StkTrInvo__CurRa__482A8B12]  DEFAULT ((0)) FOR [CurRatio]
GO

ALTER TABLE [dbo].[StkTrInvoiceDetails] ADD  CONSTRAINT [DF__StkTrInvo__Price__491EAF4B]  DEFAULT ((0)) FOR [PriceByCur]
GO

ALTER TABLE [dbo].[StkTrInvoiceDetails] ADD  CONSTRAINT [DF__StkTrInvo__QntAn__4A12D384]  DEFAULT ((0)) FOR [QntAntherUnit]
GO

ALTER TABLE [dbo].[StkTrInvoiceDetails] ADD  CONSTRAINT [DF__StkTrInvo__Price__4B06F7BD]  DEFAULT ((0)) FOR [PriceMax]
GO

ALTER TABLE [dbo].[StkTrInvoiceDetails] ADD  CONSTRAINT [DF__StkTrInvo__Price__4BFB1BF6]  DEFAULT ((0)) FOR [PriceMin]
GO




-- // procedures


GO
CREATE PROCEDURE [dbo].[StkTrInvoiceHeadList](@EmpCode int = null , @Finished bit = 0 , @Deleted bit = 0  ,@DateFrom  VARCHAR(20) = null ,@DateTo  VARCHAR(20) = null )
AS
BEGIN

SELECT  h.Serial ,  DocNo , DocDate ,h.EmpCode,
ISNULL (SUM(Qnt * Price), 0 ) TotalCash ,creator.EmpName ,AccountName,AccountCode,AccountSerial, Reserved ,  ISNULL(h.IsPosted,0) Finished ,  ISNULL(h.Deleted,0) Deleted
FROM StkTrInvoiceHead h
JOIN StkTrInvoiceDetails d 
    ON h.Serial = d.HeadSerial
JOIN Employee creator
	ON h.EmpCode = creator.EmpCode 
JOIN AccMs01 
	ON h.AccountSerial = AccMs01.Serial 
WHERE h.EmpCode = ISNULL(@Empcode , h.EmpCode)
    AND h.DocDate <= ISNULL(@DateTo , h.DocDate)
    AND h.DocDate >= ISNULL(@DateFrom , h.DocDate)
    AND h.Deleted >= ISNULL(@Deleted , h.Deleted)
    AND h.IsPosted >= ISNULL(@Finished , h.IsPosted)
group by HeadSerial ,DocNo,h.EmpCode,AccountSerial,h.Deleted, h.Serial,DocDate,EmpName,AccountCode,AccountName, Reserved,IsPosted
END


GO


CREATE PROCEDURE [dbo].[StkTrInvoiceHeadInsert]
(
    @DocNo nvarchar(16) = NULL,
	@DocDate datetime =NULL,
	@TransSerial int = 30,
	@StoreCode smallint = NULL,
	@AccountSerial int= NULL,
	@TotalCash real = NULL,
	@OrderNo int  = NULL,
	@IsPrinted bit  = NULL,
	@IsPosted bit = 0,
	@StkTr01Serial int = NULL,
	@delivery_Date datetime = NULL,
	@BuyOrSale bit = NULL,
	@ShipCode int = NULL,
	@CurCode int = NULL,
	@CurRatio money= NULL,
	@TotalByCur money =NULL,
	@EmpCode int = null,
	@AuditCode int = null,
	@Vat money = 0
	)
AS

BEGIN
    if @DocDate IS NULL
        SET @DocDate = GETDATE()
    END
    INSERT INTO   StkTrInvoiceHead
    (   DocNo,
        DocDate,
        TransSerial,
        StoreCode,
        AccountSerial,
        TotalCash,
        OrderNo,
        EmpCode,
        AuditCode ,
        IsPrinted,
        IsPosted,
        StkTr01Serial,
        delivery_Date,
        BuyOrSale,
        ShipCode,
        CurCode,
        CurRatio,
        TotalByCur ,
        Vat)
    VALUES 
        (@DocNo,
        @DocDate ,
        @TransSerial,
        @StoreCode,
        @AccountSerial,
        @TotalCash,
        @OrderNo,
        @EmpCode,
        @AuditCode,
        @IsPrinted,
        @IsPosted,
        @StkTr01Serial,
        @delivery_Date,
        @BuyOrSale,
        @ShipCode,
	    @CurCode,
        @CurRatio,
        @TotalByCur ,
        @Vat
	  )
    SELECT  SCOPE_IDENTITY() as HeadSerial
END


GO
CREATE PROCEDURE [dbo].[GetAccount](@Code  int = null , @Name   nvarchar(20) = '', @Type int )
AS
	SELECT Serial, AccountCode , AccountName 
	FROM AccMs01
	WHERE AccountName like  ('%' + @Name + '%') 
	AND AccountType = @Type
	AND AccountCode = ISNULL(@Code , AccountCode) 




GO
CREATE PROCEDURE [dbo].[StkTrInvoiceDocNo]( @TrSerial as int)
AS
SELECT ISNULL (MAX(DocNo) ,0) DocNo FROM StkTrInvoiceHead WHERE TransSerial = @TrSerial
