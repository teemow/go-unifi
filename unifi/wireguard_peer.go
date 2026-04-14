package unifi

import "context"

func (c *ApiClient) ListWireGuardPeer(ctx context.Context, site string) ([]WireGuardPeer, error) {
	return c.listWireGuardPeer(ctx, site)
}

func (c *ApiClient) GetWireGuardPeer(ctx context.Context, site, id string) (*WireGuardPeer, error) {
	return c.getWireGuardPeer(ctx, site, id)
}

func (c *ApiClient) DeleteWireGuardPeer(ctx context.Context, site, id string) error {
	return c.deleteWireGuardPeer(ctx, site, id)
}

func (c *ApiClient) CreateWireGuardPeer(
	ctx context.Context,
	site string,
	d *WireGuardPeer,
) (*WireGuardPeer, error) {
	return c.createWireGuardPeer(ctx, site, d)
}

func (c *ApiClient) UpdateWireGuardPeer(
	ctx context.Context,
	site string,
	d *WireGuardPeer,
) (*WireGuardPeer, error) {
	return c.updateWireGuardPeer(ctx, site, d)
}
