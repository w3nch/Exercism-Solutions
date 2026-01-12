"""Functions to manage a user's shopping cart items."""


def add_item(cart, items):
    """Add items to the shopping cart.

    :param cart: dict - current shopping cart
    :param items: iterable - items to add
    :return: dict - updated shopping cart
    """
    for item in items:
        cart[item] = cart.get(item, 0) + 1
    return cart


def read_notes(notes):
    """Create a shopping cart from an iterable of notes.

    :param notes: iterable - items to add
    :return: dict - shopping cart
    """
    return add_item({}, notes)


def update_recipes(ideas, updates):
    """Return an updated recipe ideas dictionary.

    :param ideas: dict - existing ideas
    :param updates: dict - recipe updates
    :return: dict - merged ideas
    """
    ideas.update(updates)
    return ideas

def sort_entries(cart):
    """Return the shopping cart sorted alphabetically by item.

    :param cart: dict - shopping cart
    :return: dict - sorted cart
    """
    return dict(sorted(cart.items()))


def send_to_store(cart, aisle_mapping):
    """Create a fulfillment cart with aisle and refrigeration info.

    :param cart: dict - shopping cart
    :param aisle_mapping: dict - item to aisle/refrigeration mapping
    :return: dict - fulfillment cart
    """
    return {
        item: [quantity, *aisle_mapping[item]]
        for item, quantity in sorted(cart.items(), reverse=True)
        if item in aisle_mapping
    }


def update_store_inventory(fulfillment_cart, store_inventory):
    """Update store inventory quantities based on a fulfillment cart.

    :param fulfillment_cart: dict - ordered items
    :param store_inventory: dict - store inventory
    :return: dict - updated inventory
    """
    for item, (ordered_qty, _, _) in fulfillment_cart.items():
        quantity, aisle, refrigerated = store_inventory[item]
        remaining = quantity - ordered_qty

        if remaining <= 0:
            store_inventory[item] = ['Out of Stock', aisle, refrigerated]
        else:
            store_inventory[item] = [remaining, aisle, refrigerated]

    return store_inventory
